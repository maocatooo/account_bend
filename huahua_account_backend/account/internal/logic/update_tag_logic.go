package logic

import (
	"context"
	"huahua_account_backend/account/internal/model"

	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.Tag) (resp *types.Tag, err error) {
	err = l.svcCtx.TagModel.Updates(l.ctx, &model.Tag{
		ID: req.ID,
	}, map[string]any{
		"Icon": req.Icon,
		"Name": req.Name,
	})
	resp = &types.Tag{
		ID:   req.ID,
		Name: req.Name,
		Icon: req.Icon,
	}
	return
}
