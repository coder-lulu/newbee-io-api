package output_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOutputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOutputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOutputTaskLogic {
	return &UpdateOutputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOutputTaskLogic) UpdateOutputTask(req *types.UpdateOutputTaskReq) (resp *types.OutputTaskInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
