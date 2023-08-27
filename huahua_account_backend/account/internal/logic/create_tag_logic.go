package logic

import (
	"context"
	"huahua_account_backend/account/internal/model"

	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTagLogic) CreateTag(req *types.Tag) (resp *types.Tag, err error) {
	uid := getUserID(l.ctx)
	tag := &model.Tag{
		Name:     req.Name,
		Uid:      uid,
		Priority: req.Priority,
		Icon:     req.Icon,
	}

	err = l.svcCtx.TagModel.Create(context.Background(), tag)
	if err != nil {
		return
	}
	resp = &types.Tag{
		ID:       tag.ID,
		Name:     tag.Name,
		Priority: tag.Priority,
		Icon:     tag.Icon,
	}
	return
}
