package field_mapping

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFieldMappingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFieldMappingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFieldMappingLogic {
	return &GetFieldMappingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFieldMappingLogic) GetFieldMapping(req *types.IDPathReq) (resp *types.FieldMappingInfoResp, err error) {
	data, err := l.svcCtx.IoRpc.GetFieldMappingById(l.ctx, &ioclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	item := transform.FieldMapping(data)
	return &types.FieldMappingInfoResp{Data: item}, nil
}
