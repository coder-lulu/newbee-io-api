package discovery_pool

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDiscoveryPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDiscoveryPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDiscoveryPoolLogic {
	return &CreateDiscoveryPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDiscoveryPoolLogic) CreateDiscoveryPool(req *types.CreateDiscoveryPoolReq) (resp *types.DiscoveryPoolInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
