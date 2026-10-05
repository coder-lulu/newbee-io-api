package field_mapping

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFieldMappingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFieldMappingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFieldMappingLogic {
	return &CreateFieldMappingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFieldMappingLogic) CreateFieldMapping(req *types.CreateFieldMappingReq) (resp *types.FieldMappingInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
