package data_target

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDataTargetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDataTargetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDataTargetLogic {
	return &UpdateDataTargetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDataTargetLogic) UpdateDataTarget(req *types.UpdateDataTargetReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
