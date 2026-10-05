package discovery_template

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDiscoveryTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDiscoveryTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDiscoveryTemplateLogic {
	return &DeleteDiscoveryTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDiscoveryTemplateLogic) DeleteDiscoveryTemplate(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
