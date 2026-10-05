package discovery_template

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiscoveryTemplateByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDiscoveryTemplateByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiscoveryTemplateByIdLogic {
	return &GetDiscoveryTemplateByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiscoveryTemplateByIdLogic) GetDiscoveryTemplateById(req *types.IDReq) (resp *types.DiscoveryTemplateInfo, err error) {
	data, err := l.svcCtx.IoRpc.GetDiscoveryTemplateById(l.ctx, &ioclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	item := transform.DiscoveryTemplate(data)
	return &item, nil
}
