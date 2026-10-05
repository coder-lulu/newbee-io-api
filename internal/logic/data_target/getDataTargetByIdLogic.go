package data_target

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataTargetByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataTargetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataTargetByIdLogic {
	return &GetDataTargetByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataTargetByIdLogic) GetDataTargetById(req *types.IDReq) (resp *types.DataTargetInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
