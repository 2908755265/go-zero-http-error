package go_zero_http_error

import (
	"errors"
	"fmt"
)

type httpError struct {
	e error
	c int
	d any
}

type errResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type CodeError interface {
	CodeErr() (int, any)
}

func (e *httpError) Error() string {
	return e.e.Error()
}

func (e *httpError) CodeErr() (int, any) {
	return e.c, errResponse{
		Code: e.c,
		Msg:  e.Error(),
		Data: e.d,
	}
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
