package go_zero_http_error

import (
	"encoding/json"
	"errors"
	"fmt"
)

type httpError struct {
	e error
	c int
	d any
}

type errResponse struct {
	Code int
	Msg  string
	Data any
}

type CodeError interface {
	CodeErr() (int, any)
}

func (e *httpError) Error() string {
	return e.e.Error()
}

func (e *httpError) CodeErr() (int, any) {
	body, _ := json.Marshal(errResponse{
		Code: e.c,
		Msg:  e.Error(),
		Data: e.d,
	})
	return e.c, body
}

func NewError(code int, data any, err error) error {
	if err == nil {
		err = errors.New(fmt.Sprintf("err %d", code))
	}
	return &httpError{
		e: err,
		c: code,
		d: data,
	}
}
