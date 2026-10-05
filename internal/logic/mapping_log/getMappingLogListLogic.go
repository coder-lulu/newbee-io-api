package mapping_log

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMappingLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMappingLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMappingLogListLogic {
	return &GetMappingLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMappingLogListLogic) GetMappingLogList(req *types.GetMappingLogListReq) (resp *types.MappingLogListResp, err error) {
	data, err := l.svcCtx.IoRpc.GetMappingLogList(l.ctx, &ioclient.MappingLogListReq{Page: req.Page, PageSize: req.PageSize, FieldMappingId: req.FieldMappingId})
	if err != nil {
		return nil, err
	}
	resp = &types.MappingLogListResp{}
	resp.Data.Total = data.Total
	resp.Data.Data = make([]types.MappingLogInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data.Data = append(resp.Data.Data, transform.MappingLog(item))
	}
	return resp, nil
}
