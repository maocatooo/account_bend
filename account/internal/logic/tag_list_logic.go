package logic

import (
	"context"

	"account_bend/account/internal/svc"
	"account_bend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagListLogic {
	return &TagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagListLogic) TagList() (resp []*types.Tag, err error) {
	uid := getUserID(l.ctx)

	res, err := l.svcCtx.TagModel.FindByUserID(l.ctx, uid)
	if err != nil {
		return nil, err
	}

	resp = make([]*types.Tag, len(res))
	for i, v := range res {
		resp[i] = &types.Tag{
			ID:          v.ID,
			Name:        v.Name,
			CreatedTime: int(v.CreatedAt.UnixMilli()),
			Priority:    v.Priority,
			Icon:        v.Icon,
		}
	}
	return
}
