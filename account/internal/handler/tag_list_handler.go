package handler

import (
	"net/http"

	"account_bend/account/internal/logic"
	"account_bend/account/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func tagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTagListLogic(r.Context(), svcCtx)
		resp, err := l.TagList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
