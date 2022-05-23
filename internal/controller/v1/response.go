package v1

import (
	"encoding/json"
	"net/http"
)

type response struct {
	http.ResponseWriter
	headers    map[string]string
	statusCode int
}

func newResponse(w http.ResponseWriter) *response {
	return &response{
		w,
		make(map[string]string),
		defaultStatusCode,
	}
}

func (r *response) withStatusCode(sc int) *response {
	r.statusCode = sc
	return r
}

func (r *response) withContentType(ct string) *response {
	r.headers["Content-Type"] = ct
	return r
}

func (r *response) build() *response {
	for key, value := range r.headers {
		r.Header().Set(key, value)
	}
	r.WriteHeader(r.statusCode)
	return r
}

func (r *response) withBytes(data []byte) *response {
	_, err := r.build().Write(data)
	if err != nil {
		panic(err)
	}
	return r
}

func (r *response) withJSON(elem interface{}) *response {
	jsonResp, err := json.Marshal(elem)
	if err != nil {
		panic(err)
	}
	return r.
		withContentType("application/json").
		withBytes(jsonResp)
}

func (r *response) withString(str string) *response {
	return r.
		withContentType("text/plain").
		withBytes([]byte(str))
}

func (r *response) withHTTPError(sc int, msg string, errs ...*Error) *response {
	err := NewError(sc, msg, errs...)
	return r.
		withStatusCode(sc).
		withJSON(err)
}
