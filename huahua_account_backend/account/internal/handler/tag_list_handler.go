package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"huahua_account_backend/account/internal/logic"
	"huahua_account_backend/account/internal/svc"
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
