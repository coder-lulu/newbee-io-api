package output_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOutputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOutputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOutputTaskLogic {
	return &GetOutputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOutputTaskLogic) GetOutputTask(req *types.IDPathReq) (resp *types.OutputTaskInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
