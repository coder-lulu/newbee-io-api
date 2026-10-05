package discovery_pool

import (
	"context"

	"github.com/coder-lulu/newbee-common/v2/i18n"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/types/io"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiscoveryPoolListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDiscoveryPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiscoveryPoolListLogic {
	return &GetDiscoveryPoolListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiscoveryPoolListLogic) GetDiscoveryPoolList(req *types.GetDiscoveryPoolListReq) (resp *types.DiscoveryPoolListResp, err error) {
	// 调用RPC服务
	rpcResp, err := l.svcCtx.IoRpc.GetDiscoveryPoolList(l.ctx, &io.DiscoveryPoolListReq{
		Page:           req.Page,
		PageSize:       req.PageSize,
		Name:           req.Name,
		DiscoveryType:  req.DiscoveryType,
		PoolStatus:     req.PoolStatus,
		ApprovalStatus: req.ApprovalStatus,
	})
	if err != nil {
		return nil, err
	}

	// 转换RPC响应为API响应
	resp = &types.DiscoveryPoolListResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.DiscoveryPoolListData{
			BaseListInfo: types.BaseListInfo{
				Total: rpcResp.Total,
			},
			Data: make([]types.DiscoveryPoolInfo, 0, len(rpcResp.Data)),
		},
	}

	// 转换数据列表
	for _, v := range rpcResp.Data {
		info := types.DiscoveryPoolInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Name:            safeString(v.Name),
			Description:     safeString(v.Description),
			DiscoveryType:   safeString(v.DiscoveryType),
			PoolStatus:      safeString(v.PoolStatus),
			DiscoveryConfig: safeString(v.DiscoveryConfig),
			Schedule:        safeString(v.Schedule),
			BatchSize:       safeInt(v.BatchSize),
			ConcurrentLimit: safeInt(v.ConcurrentLimit),
			MaxRetry:        safeInt(v.MaxRetry),
			RetryInterval:   safeInt(v.RetryInterval),
			FieldMapping:    safeString(v.FieldMapping),
			TotalRuns:       safeInt64(v.TotalRuns),
			SuccessRuns:     safeInt64(v.SuccessRuns),
			FailedRuns:      safeInt64(v.FailedRuns),
		}

		// 可选字段
		if v.LastRunAt != nil {
			info.LastRunAt = v.LastRunAt
		}
		if v.LastSuccessAt != nil {
			info.LastSuccessAt = v.LastSuccessAt
		}
		if v.LastError != nil {
			info.LastError = v.LastError
		}
		if v.ApprovalStatus != nil {
			info.ApprovalStatus = *v.ApprovalStatus
		}
		if v.ApprovedBy != nil {
			info.ApprovedBy = v.ApprovedBy
		}
		if v.ApprovedAt != nil {
			info.ApprovedAt = v.ApprovedAt
		}
		if v.RejectionReason != nil {
			info.RejectionReason = v.RejectionReason
		}
		if v.Metadata != nil {
			info.Metadata = *v.Metadata
		}

		resp.Data.Data = append(resp.Data.Data, info)
	}

	return resp, nil
}

// 辅助函数：安全地解引用字符串指针
func safeString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// 辅助函数：安全地解引用int64指针并转换为int
func safeInt(i *int64) int {
	if i != nil {
		return int(*i)
	}
	return 0
}

// 辅助函数：安全地解引用int64指针
func safeInt64(i *int64) int64 {
	if i != nil {
		return *i
	}
	return 0
}
