package input_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

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
	data, err := l.svcCtx.IoRpc.GetInputTaskList(l.ctx, &ioclient.InputTaskListReq{Page: req.Page, PageSize: req.PageSize, TaskName: req.Name, InputSource: req.TaskType, TaskStatus: req.TaskStatus, DiscoveryPoolId: req.DiscoveryPoolId})
	if err != nil {
		return nil, err
	}
	resp = &types.InputTaskListResp{}
	resp.Data.Total = data.Total
	resp.Data.Data = make([]types.InputTaskInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data.Data = append(resp.Data.Data, transform.InputTask(item))
	}
	return resp, nil
}
