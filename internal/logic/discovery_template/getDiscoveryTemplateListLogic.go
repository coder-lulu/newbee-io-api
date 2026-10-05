package discovery_template

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiscoveryTemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDiscoveryTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiscoveryTemplateListLogic {
	return &GetDiscoveryTemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiscoveryTemplateListLogic) GetDiscoveryTemplateList(req *types.DiscoveryTemplateListReq) (resp *types.DiscoveryTemplateListResp, err error) {
	data, err := l.svcCtx.IoRpc.GetDiscoveryTemplateList(l.ctx, &ioclient.DiscoveryTemplateListReq{Page: req.Page, PageSize: req.PageSize, TemplateType: req.TemplateType, IsPublic: req.IsPublic, IsSystem: req.IsSystem, Keyword: req.Keyword})
	if err != nil {
		return nil, err
	}
	resp = &types.DiscoveryTemplateListResp{}
	resp.Total = data.Total
	resp.Data = make([]types.DiscoveryTemplateInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data = append(resp.Data, transform.DiscoveryTemplate(item))
	}
	return resp, nil
}
