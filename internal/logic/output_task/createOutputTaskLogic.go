package output_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOutputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOutputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOutputTaskLogic {
	return &CreateOutputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOutputTaskLogic) CreateOutputTask(req *types.CreateOutputTaskReq) (resp *types.OutputTaskInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
