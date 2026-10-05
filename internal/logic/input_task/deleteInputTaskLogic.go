package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInputTaskLogic {
	return &DeleteInputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInputTaskLogic) DeleteInputTask(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
