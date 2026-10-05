package data_target

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDataTargetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDataTargetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDataTargetLogic {
	return &CreateDataTargetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDataTargetLogic) CreateDataTarget(req *types.CreateDataTargetReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
