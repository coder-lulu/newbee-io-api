package task_log

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogLogic {
	return &CreateTaskLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogLogic) CreateTaskLog(req *types.CreateTaskLogReq) (resp *types.TaskLogInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
