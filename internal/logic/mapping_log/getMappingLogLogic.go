package mapping_log

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMappingLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMappingLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMappingLogLogic {
	return &GetMappingLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMappingLogLogic) GetMappingLog(req *types.IDPathReq) (resp *types.MappingLogInfoResp, err error) {
	data, err := l.svcCtx.IoRpc.GetMappingLogById(l.ctx, &ioclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	item := transform.MappingLog(data)
	return &types.MappingLogInfoResp{Data: item}, nil
}
