package mapping_log

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMappingLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMappingLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMappingLogListLogic {
	return &GetMappingLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMappingLogListLogic) GetMappingLogList(req *types.GetMappingLogListReq) (resp *types.MappingLogListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
