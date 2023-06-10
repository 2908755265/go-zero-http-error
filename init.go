package go_zero_http_error

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func init() {
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		ce, ok := err.(CodeError)
		if ok {
			return ce.CodeErr()
		}
		body, _ := json.Marshal(errResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
			Data: nil,
		})
		return http.StatusInternalServerError, body
	})
}
