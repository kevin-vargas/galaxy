package v1

import "net/http"

type StatusHTTP int

func (e StatusHTTP) toError(msg string) *Error {
	return &Error{
		Code:    http.StatusText(int(e)),
		Message: msg,
	}
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Errors []*Error

type ErrorHTTP struct {
	*Error
	Errors Errors `json:"errors"`
}

func NewError(e int, msg string, errs ...*Error) *ErrorHTTP {
	return &ErrorHTTP{
		StatusHTTP(e).toError(msg),
		errs,
	}
}

func IncorrectHeader(header, msg string) *Error {
	return &Error{
		"incorrect header " + header,
		msg,
	}
}

func IncorrectQueryParameter(qp, msg string) *Error {
	return &Error{
		"incorrect parameter " + qp,
		"Query Param " + qp + " " + msg,
	}
}

func IncorrectBody(msg string) *Error {
	return &Error{
		"incorrect body",
		msg,
	}
}
