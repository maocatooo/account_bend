package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"huahua_account_backend/account/internal/logic"
	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"
)

func deleteJournalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Journal
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteJournalLogic(r.Context(), svcCtx)
		err := l.DeleteJournal(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
