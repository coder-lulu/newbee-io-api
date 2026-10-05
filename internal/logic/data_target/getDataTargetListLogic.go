package data_target

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataTargetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataTargetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataTargetListLogic {
	return &GetDataTargetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataTargetListLogic) GetDataTargetList(req *types.DataTargetListReq) (resp *types.DataTargetListResp, err error) {
	data, err := l.svcCtx.IoRpc.GetDataTargetList(l.ctx, &ioclient.DataTargetListReq{Page: req.Page, PageSize: req.PageSize, TargetType: req.TargetType, TargetName: req.Keyword})
	if err != nil {
		return nil, err
	}
	resp = &types.DataTargetListResp{}
	resp.Total = data.Total
	resp.Data = make([]types.DataTargetInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data = append(resp.Data, transform.DataTarget(item))
	}
	return resp, nil
}
