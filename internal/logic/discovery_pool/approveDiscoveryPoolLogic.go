package discovery_pool

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveDiscoveryPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApproveDiscoveryPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveDiscoveryPoolLogic {
	return &ApproveDiscoveryPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApproveDiscoveryPoolLogic) ApproveDiscoveryPool(req *types.ApproveDiscoveryPoolReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
