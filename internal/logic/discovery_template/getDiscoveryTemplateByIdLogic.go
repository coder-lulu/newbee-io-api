package discovery_template

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
