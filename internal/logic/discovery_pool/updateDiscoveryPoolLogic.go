package discovery_pool

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDiscoveryPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDiscoveryPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDiscoveryPoolLogic {
	return &UpdateDiscoveryPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDiscoveryPoolLogic) UpdateDiscoveryPool(req *types.UpdateDiscoveryPoolReq) (resp *types.DiscoveryPoolInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
