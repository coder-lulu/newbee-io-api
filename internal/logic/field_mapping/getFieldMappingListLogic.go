package field_mapping

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFieldMappingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFieldMappingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFieldMappingListLogic {
	return &GetFieldMappingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFieldMappingListLogic) GetFieldMappingList(req *types.GetFieldMappingListReq) (resp *types.FieldMappingListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
