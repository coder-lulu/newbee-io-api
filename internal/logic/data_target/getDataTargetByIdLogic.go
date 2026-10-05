package data_target

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

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
	data, err := l.svcCtx.IoRpc.GetDataTargetById(l.ctx, &ioclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	item := transform.DataTarget(data)
	return &item, nil
}
