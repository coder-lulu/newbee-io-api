package output_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOutputTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOutputTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOutputTaskLogic {
	return &GetOutputTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOutputTaskLogic) GetOutputTask(req *types.IDPathReq) (resp *types.OutputTaskInfoResp, err error) {
	data, err := l.svcCtx.IoRpc.GetOutputTaskById(l.ctx, &ioclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	item := transform.OutputTask(data)
	return &types.OutputTaskInfoResp{Data: item}, nil
}
