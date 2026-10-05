package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

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
	data, err := l.svcCtx.IoRpc.GetInputTaskById(l.ctx, &ioclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	item := transform.InputTask(data)
	return &types.InputTaskInfoResp{Data: item}, nil
}
