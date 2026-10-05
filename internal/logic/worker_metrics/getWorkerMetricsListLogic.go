package worker_metrics

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
