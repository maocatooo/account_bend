package logic

import (
	"context"

	"account_bend/account/internal/svc"
	"account_bend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NamePromptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamePromptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NamePromptLogic {
	return &NamePromptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamePromptLogic) NamePrompt(req *types.Tid) (resp []string, err error) {
	uid := getUserID(l.ctx)
	resp, err = l.svcCtx.BookJournalModel.NamePrompt(l.ctx, uid, req.Tid)
	return
}
