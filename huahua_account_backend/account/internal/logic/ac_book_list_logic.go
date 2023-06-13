package logic

import (
	"context"
	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcBookListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAcBookListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcBookListLogic {
	return &AcBookListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AcBookListLogic) AcBookList() (resp []*types.AcBook, err error) {

	uid := getUserID(l.ctx)

	res, err := l.svcCtx.BookModel.FindByUid(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	for _, item := range res {
		resp = append(resp, &types.AcBook{
			Id:          item.ID,
			Name:        item.Name,
			CreatedTime: int(item.CreatedAt.UnixMilli()),
			Tp:          item.Tp,
			Uid:         item.Uid,
		})
	}
	return
}
