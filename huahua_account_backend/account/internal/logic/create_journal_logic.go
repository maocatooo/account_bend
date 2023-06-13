package logic

import (
	"context"
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

	journal := &model.AcBookJournal{
		Amount:   req.Amount,
		Date:     time.Unix(int64(req.Date), 0),
		Tid:      req.Tid,
		Tname:    req.Tname,
		Notes:    req.Notes,
		AcBookId: req.AcBookId,
		Uid:      uid,
	}
	err = l.svcCtx.AcBookJournalModel.Create(context.Background(), journal)
	if err != nil {
		return nil, err
	}
	resp = &types.Journal{
		Id:       journal.Id,
		Amount:   journal.Amount,
		Date:     int(journal.Date.Unix()),
		Tid:      journal.Tid,
		Tname:    journal.Tname,
		Notes:    journal.Notes,
		AcBookId: journal.AcBookId,
		Uid:      journal.Uid,
	}

	return
}
