package task_log

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogListLogic {
	return &GetTaskLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskLogListLogic) GetTaskLogList(req *types.GetTaskLogListReq) (resp *types.TaskLogListResp, err error) {
	data, err := l.svcCtx.IoRpc.GetTaskLogList(l.ctx, &ioclient.TaskLogListReq{Page: req.Page, PageSize: req.PageSize, TaskId: req.TaskId, TaskType: req.TaskType, LogLevel: req.Level})
	if err != nil {
		return nil, err
	}
	resp = &types.TaskLogListResp{}
	resp.Data.Total = data.Total
	resp.Data.Data = make([]types.TaskLogInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data.Data = append(resp.Data.Data, transform.TaskLog(item))
	}
	return resp, nil
}
