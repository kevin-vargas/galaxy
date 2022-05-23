package v1

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevin-vargas/galaxy/internal/usecase"
	"github.com/kevin-vargas/galaxy/internal/usecase/finder"
	"github.com/kevin-vargas/galaxy/internal/usecase/mocks"
	"github.com/kevin-vargas/galaxy/internal/usecase/satellites"
	"github.com/stretchr/testify/assert"
)

func Test_Topsecret(t *testing.T) {
	testCases := []struct {
		desc       string
		method     string
		statusCode int
		body       []byte
		url        string
		expect     string
		header     http.Header
	}{
		{
			"With Out Content Type Application Json",
			http.MethodPost,
			415,
			nil,
			"/topsecret",
			expect_with_out_content_type,
			http.Header{},
		},
		{
			"With Incorrect Body",
			http.MethodPost,
			400,
			nil,
			"/topsecret",
			expect_with_out_body,
			http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			"With Incorrect Body Field",
			http.MethodPost,
			400,
			[]byte(`{"satellites":[{"name":1,"distance":200.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":500,"message":["","es","","","secreto"]},{"name":"sato","distance":904.3973,"message":["este","","un","",""]}]}`),
			"/topsecret",
			expect_with_incorrect_field_body,
			http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			"With Unknown Satellite",
			http.MethodPost,
			409,
			[]byte(`{"satellites":[{"name":"test","distance":200.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":500,"message":["","es","","","secreto"]},{"name":"sato","distance":904.3973,"message":["este","","un","",""]}]}`),
			"/topsecret",
			expect_with_unknown_satellite,
			http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			"With Correct Body",
			http.MethodPost,
			200,
			[]byte(`{"satellites":[{"name":"kenobi","distance":200.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":500,"message":["","es","","","secreto"]},{"name":"sato","distance":904.3973,"message":["este","","un","",""]}]}`),
			"/topsecret",
			expect_with_correct_body,
			http.Header{"Content-Type": []string{"application/json"}},
		},
	}
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("On_%s_%s", tt.url, tt.desc), func(t *testing.T) {
			// arrange
			s := satellites.New()
			f := finder.New()
			triangulation := usecase.NewTriangulation(s, f)
			// triangulation := mocks.Triangulation{}
			session := mocks.Session{}

			handler := newRoutes(triangulation, &session)
			req, res := getReqRes(tt.method, tt.body, tt.url, tt.header)

			// act
			handler.ServeHTTP(res, req)
			resultStatusCode := res.Result().StatusCode
			resultBody := string(res.Body.Bytes())

			// assert
			assert.Equal(t, tt.statusCode, resultStatusCode)
			if resultBody != "" {
				assert.JSONEq(t, tt.expect, resultBody)
			}

		})
	}
}

var getReqRes = func(method string, body []byte, url string, headers http.Header) (*http.Request, *httptest.ResponseRecorder) {
	reader := bytes.NewReader(body)
	req := httptest.NewRequest(method, url, reader)
	req.Header = headers
	res := httptest.NewRecorder()
	return req, res
}

const (
	expect_with_correct_body         = `{"positions":[{"x":-395.04094970246405,"y":-29.754301785216}],"message":"este es un mensaje secreto"}`
	expect_with_incorrect_field_body = `{"code":"Bad Request","message":"On Request","errors":[{"code":"incorrect body","message":"wrong Type provided for field satellites.name"}]}`
	expect_with_out_content_type     = `{"code":"Unsupported Media Type","message":"On Request","errors":[{"code":"incorrect header Content-Type","message":"Content Type is not application/json"}]}`
	expect_with_out_body             = `{"code":"Bad Request","message":"On Request","errors":[{"code":"incorrect body","message":"EOF"}]}`
	expect_with_unknown_satellite    = `{"code":"Conflict","message":"get positions","errors":[{"code":"incorrect body","message":"unknown satellite"}]}`
	audit_expect_with_out_body       = `{"code":"Bad Request","message":"On audit request","errors":[{"code":"incorrect body","message":"EOF"}]}`
)
