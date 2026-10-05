package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateInputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInputTaskLogic {
	return &CreateInputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateInputTaskLogic) CreateInputTask(req *types.CreateInputTaskReq) (resp *types.InputTaskInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
