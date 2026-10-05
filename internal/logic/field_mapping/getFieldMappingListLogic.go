package field_mapping

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFieldMappingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFieldMappingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFieldMappingListLogic {
	return &GetFieldMappingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFieldMappingListLogic) GetFieldMappingList(req *types.GetFieldMappingListReq) (resp *types.FieldMappingListResp, err error) {
	data, err := l.svcCtx.IoRpc.GetFieldMappingList(l.ctx, &ioclient.FieldMappingListReq{Page: req.Page, PageSize: req.PageSize, MappingName: req.MappingName, MappingType: req.MappingType, IsActive: req.IsActive})
	if err != nil {
		return nil, err
	}
	resp = &types.FieldMappingListResp{}
	resp.Data.Total = data.Total
	resp.Data.Data = make([]types.FieldMappingInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data.Data = append(resp.Data.Data, transform.FieldMapping(item))
	}
	return resp, nil
}
