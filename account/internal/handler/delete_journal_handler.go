package handler

import (
	"net/http"

	"account_bend/account/internal/logic"
	"account_bend/account/internal/svc"
	"account_bend/account/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
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
