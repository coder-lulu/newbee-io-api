package output_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOutputTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOutputTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOutputTaskListLogic {
	return &GetOutputTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOutputTaskListLogic) GetOutputTaskList(req *types.GetOutputTaskListReq) (resp *types.OutputTaskListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
