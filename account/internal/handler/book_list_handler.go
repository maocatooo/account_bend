package handler

import (
	"net/http"

	"account_bend/account/internal/logic"
	"account_bend/account/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BookListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewBookListLogic(r.Context(), svcCtx)
		resp, err := l.BookList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
