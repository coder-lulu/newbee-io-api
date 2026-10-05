package discovery_template

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDiscoveryTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDiscoveryTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDiscoveryTemplateLogic {
	return &UpdateDiscoveryTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDiscoveryTemplateLogic) UpdateDiscoveryTemplate(req *types.UpdateDiscoveryTemplateReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
