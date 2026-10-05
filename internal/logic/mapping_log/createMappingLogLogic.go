package mapping_log

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMappingLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMappingLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMappingLogLogic {
	return &CreateMappingLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMappingLogLogic) CreateMappingLog(req *types.CreateMappingLogReq) (resp *types.MappingLogInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
