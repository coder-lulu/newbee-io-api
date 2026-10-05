package worker_metrics

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/transform"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkerMetricsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWorkerMetricsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkerMetricsListLogic {
	return &GetWorkerMetricsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWorkerMetricsListLogic) GetWorkerMetricsList(req *types.WorkerMetricsListReq) (resp *types.WorkerMetricsListResp, err error) {
	data, err := l.svcCtx.IoRpc.GetWorkerMetricsList(l.ctx, &ioclient.WorkerMetricsListReq{Page: req.Page, PageSize: req.PageSize, WorkerId: req.WorkerId})
	if err != nil {
		return nil, err
	}
	resp = &types.WorkerMetricsListResp{}
	resp.Total = data.Total
	resp.Data = make([]types.WorkerMetricsInfo, 0, len(data.Data))
	for _, item := range data.Data {
		resp.Data = append(resp.Data, transform.WorkerMetrics(item))
	}
	return resp, nil
}
