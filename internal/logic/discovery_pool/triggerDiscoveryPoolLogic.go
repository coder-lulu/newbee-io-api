package discovery_pool

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TriggerDiscoveryPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTriggerDiscoveryPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TriggerDiscoveryPoolLogic {
	return &TriggerDiscoveryPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TriggerDiscoveryPoolLogic) TriggerDiscoveryPool(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
