package logic

import (
	"account_bend/account/internal/svc"
	"account_bend/account/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type BookListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BookListLogic {
	return &BookListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookListLogic) BookList() (resp []*types.Book, err error) {

	uid := getUserID(l.ctx)

	res, err := l.svcCtx.BookModel.FindByUid(l.ctx, uid)
	if err != nil {
		return nil, err
	}
	for _, item := range res {
		resp = append(resp, &types.Book{
			ID:          item.ID,
			Name:        item.Name,
			CreatedTime: int(item.CreatedAt.UnixMilli()),
			Tp:          item.Tp,
			Uid:         item.Uid,
		})
	}
	return
}
