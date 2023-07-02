package logic

import (
	"context"
	"fmt"
	"huahua_account_backend/account/internal/model"
	"time"

	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateJournalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateJournalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateJournalLogic {
	return &CreateJournalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateJournalLogic) CreateJournal(req *types.Journal) (resp *types.Journal, err error) {

	uid := getUserID(l.ctx)

	journal := &model.BookJournal{
		Amount: req.Amount,
		Date:   time.UnixMilli(int64(req.Date)),
		Tid:    req.Tid,
		Tname:  req.Tname,
		Record: req.Record,
		BookID: req.BookID,
		Uid:    uid,
	}
	fmt.Println(journal)
	err = l.svcCtx.BookJournalModel.Create(context.Background(), journal)
	if err != nil {
		return nil, err
	}
	resp = &types.Journal{
		ID:     journal.ID,
		Amount: journal.Amount,
		Date:   int(journal.Date.Unix()),
		Tid:    journal.Tid,
		Tname:  journal.Tname,
		Record: journal.Record,
		BookID: journal.BookID,
		Uid:    journal.Uid,
	}

	return
}
