// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"account_bend/account/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/book",
				Handler: BookListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/tag",
				Handler: tagListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/name_prompt",
				Handler: NamePromptHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/tag",
				Handler: createTagHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/tag",
				Handler: updateTagHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/journal",
				Handler: createJournalHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/journal",
				Handler: updateJournalHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/journal",
				Handler: deleteJournalHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/journal",
				Handler: journalListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}