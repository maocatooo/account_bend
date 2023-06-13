package logic

import (
	"context"

	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JournalListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJournalListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JournalListLogic {
	return &JournalListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JournalListLogic) JournalList(req *types.Journal) (resp []*types.Journal, err error) {

	journals, err := l.svcCtx.AcBookJournalModel.FindByBookId(context.Background(), req.AcBookId)
	if err != nil {
		return
	}

	for _, journal := range journals {
		resp = append(resp, &types.Journal{
			Id:       journal.Id,
			Amount:   journal.Amount,
			Date:     int(journal.Date.Unix()),
			Tid:      journal.Tid,
			Tname:    journal.Tname,
			Notes:    journal.Notes,
			AcBookId: journal.AcBookId,
			Uid:      journal.Uid,
		})
	}

	return
}
