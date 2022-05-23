package ctx

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	// arrange
	toTest := contextKeyTraceID

	// act
	result := toTest.String()

	// assert
	assert.Equal(t, "request_trace_id", result)
}

func Test_Trace_ID(t *testing.T) {
	// arrange
	ctx := context.Background()
	id := "fake_id"
	ctx = context.WithValue(ctx, contextKeyTraceID, id)

	// act
	result, ok := GetTraceID(ctx)

	// assert
	assert.Equal(t, result, id)
	assert.True(t, ok)
}

func Test_Request_Context_Build(t *testing.T) {

	testCases := []struct {
		desc   string
		header string
		value  string
		expect string
	}{
		{
			desc:   "with X-Trace-ID",
			header: "X-Trace-ID",
			value:  "fake_id",
			expect: "fake_id",
		},
		{
			desc:   "with x-trace-id",
			header: "x-trace-id",
			value:  "fake_id",
			expect: "fake_id",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			// arrange
			ctx := context.Background()
			r := getReq(tt.value, tt.header)
			// act
			result := RequestContextBuild(ctx, r)
			id, ok := GetTraceID(result)
			// assert
			assert.Equal(t, id, tt.expect)
			assert.True(t, ok)
		})
	}

}

var getReq = func(id string, header string) *http.Request {
	reader := bytes.NewReader([]byte(`sample`))
	req := httptest.NewRequest(http.MethodGet, "https://www.some-domain.com", reader)
	req.Header.Add(header, id)
	return req
}
