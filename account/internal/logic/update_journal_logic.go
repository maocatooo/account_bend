package logic

import (
	"account_bend/account/internal/model"
	"context"
	"time"

	"account_bend/account/internal/svc"
	"account_bend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJournalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateJournalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateJournalLogic {
	return &UpdateJournalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateJournalLogic) UpdateJournal(req *types.Journal) (resp *types.Journal, err error) {
	uid := getUserID(l.ctx)

	journal := &model.BookJournal{
		ID:     req.ID,
		Amount: req.Amount,
		Name:   req.Name,
		Date:   time.UnixMilli(int64(req.Date)),
		Tid:    req.Tid,
		Tname:  req.Tname,
		Record: req.Record,
		BookID: req.BookID,
		Uid:    uid,
	}

	err = l.svcCtx.BookJournalModel.Update(context.Background(), journal)
	if err != nil {
		return nil, err
	}
	resp = &types.Journal{
		ID:     journal.ID,
		Amount: journal.Amount,
		Name:   journal.Name,
		Date:   int(journal.Date.UnixMilli()),
		Tid:    journal.Tid,
		Tname:  journal.Tname,
		Record: journal.Record,
		BookID: journal.BookID,
		Uid:    journal.Uid,
	}
	return
}
