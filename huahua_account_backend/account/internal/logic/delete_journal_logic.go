package logic

import (
	"context"

	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteJournalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteJournalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteJournalLogic {
	return &DeleteJournalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteJournalLogic) DeleteJournal(req *types.Journal) error {

	err := l.svcCtx.BookJournalModel.DeleteById(l.ctx, req.ID)
	if err != nil {
		return err
	}
	return nil
}
