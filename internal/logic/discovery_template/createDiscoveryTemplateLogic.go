package discovery_template

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDiscoveryTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDiscoveryTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDiscoveryTemplateLogic {
	return &CreateDiscoveryTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDiscoveryTemplateLogic) CreateDiscoveryTemplate(req *types.CreateDiscoveryTemplateReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
