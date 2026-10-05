package output_task

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

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
	data, err := l.svcCtx.IoRpc.GetOutputTaskList(l.ctx, &ioclient.OutputTaskListReq{Page: req.Page, PageSize: req.PageSize, TaskName: req.Name, OutputTarget: req.OutputType, TaskStatus: req.TaskStatus})
	if err != nil {
		return nil, err
	}
	resp = &types.OutputTaskListResp{}
	resp.Data.Total = data.Total
	resp.Data.Data = make([]types.OutputTaskInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data.Data = append(resp.Data.Data, transform.OutputTask(item))
	}
	return resp, nil
}
