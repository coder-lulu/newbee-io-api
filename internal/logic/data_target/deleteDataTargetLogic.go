package data_target

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDataTargetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDataTargetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDataTargetLogic {
	return &DeleteDataTargetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDataTargetLogic) DeleteDataTarget(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
