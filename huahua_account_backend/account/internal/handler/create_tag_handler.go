package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"huahua_account_backend/account/internal/logic"
	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"
)

func createTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Tag
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateTagLogic(r.Context(), svcCtx)
		resp, err := l.CreateTag(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
