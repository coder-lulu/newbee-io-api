package task_log

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogLogic {
	return &GetTaskLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskLogLogic) GetTaskLog(req *types.IDPathReq) (resp *types.TaskLogInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
