package main

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func init() {
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		res := map[string]any{}
		switch e := err.(type) {
		case error:
			res["msg"] = e.Error()
			return http.StatusOK, res
		default:
			return http.StatusInternalServerError, res
		}
	})
}
