package field_mapping

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFieldMappingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFieldMappingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFieldMappingLogic {
	return &UpdateFieldMappingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFieldMappingLogic) UpdateFieldMapping(req *types.UpdateFieldMappingReq) (resp *types.FieldMappingInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
