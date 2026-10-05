package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveInputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApproveInputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveInputTaskLogic {
	return &ApproveInputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApproveInputTaskLogic) ApproveInputTask(req *types.ApproveInputTaskReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
