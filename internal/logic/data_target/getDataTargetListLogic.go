package data_target

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataTargetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataTargetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataTargetListLogic {
	return &GetDataTargetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataTargetListLogic) GetDataTargetList(req *types.DataTargetListReq) (resp *types.DataTargetListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
