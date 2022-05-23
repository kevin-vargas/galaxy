package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kevin-vargas/galaxy/infra/ctx"
)

func Test_WithCustomContext(t *testing.T) {
	// arrange
	req, res := getReqRes(http.MethodGet)
	var id string
	var ok bool
	var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ctx := r.Context()
		id, ok = ctx.GetTraceID(_ctx)
		_, err := w.Write([]byte("bar"))
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	// act
	WithCustomContext(testHandler).ServeHTTP(res, req)

	// assert
	assert.Equal(t, "trace_id", id)
	assert.True(t, ok)
}

var getReqRes = func(method string) (*http.Request, *httptest.ResponseRecorder) {
	reader := bytes.NewReader([]byte(`sample`))
	req := httptest.NewRequest(method, "https://www.some-domain.com", reader)
	req.Header.Add("X-Trace-ID", "trace_id")
	req.AddCookie(&http.Cookie{
		Name:  "session_token",
		Value: "test",
	})
	res := httptest.NewRecorder()
	return req, res
}
