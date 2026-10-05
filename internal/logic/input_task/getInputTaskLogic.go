package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInputTaskLogic {
	return &GetInputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInputTaskLogic) GetInputTask(req *types.IDPathReq) (resp *types.InputTaskInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
