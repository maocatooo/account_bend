package logic

import (
	"context"

	"account_bend/account/internal/svc"
	"account_bend/account/internal/types"

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

	journals, err := l.svcCtx.BookJournalModel.FindByTypes(context.Background(), req)
	if err != nil {
		return
	}

	for _, journal := range journals {
		resp = append(resp, &types.Journal{
			ID:     journal.ID,
			Amount: journal.Amount,
			Name:   journal.Name,
			Date:   int(journal.Date.UnixMilli()),
			Tid:    journal.Tid,
			Tname:  journal.Tname,
			Record: journal.Record,
			BookID: journal.BookID,
			Uid:    journal.Uid,
		})
	}

	return
}
