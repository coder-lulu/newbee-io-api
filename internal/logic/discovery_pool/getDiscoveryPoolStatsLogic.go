package discovery_pool

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiscoveryPoolStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDiscoveryPoolStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiscoveryPoolStatsLogic {
	return &GetDiscoveryPoolStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiscoveryPoolStatsLogic) GetDiscoveryPoolStats(req *types.IDPathReq) (resp *types.DiscoveryPoolInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
