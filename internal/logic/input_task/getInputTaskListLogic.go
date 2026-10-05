package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInputTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInputTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInputTaskListLogic {
	return &GetInputTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInputTaskListLogic) GetInputTaskList(req *types.GetInputTaskListReq) (resp *types.InputTaskListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
