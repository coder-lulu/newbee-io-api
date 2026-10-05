package worker_metrics

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkerMetricsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWorkerMetricsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkerMetricsByIdLogic {
	return &GetWorkerMetricsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWorkerMetricsByIdLogic) GetWorkerMetricsById(req *types.IDReq) (resp *types.WorkerMetricsInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
