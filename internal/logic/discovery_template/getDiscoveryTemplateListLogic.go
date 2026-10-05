package discovery_template

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
