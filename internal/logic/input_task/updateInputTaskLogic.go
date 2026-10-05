package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInputTaskLogic {
	return &UpdateInputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInputTaskLogic) UpdateInputTask(req *types.UpdateInputTaskReq) (resp *types.InputTaskInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
