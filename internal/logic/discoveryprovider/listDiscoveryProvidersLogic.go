package discoveryprovider

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/types/io"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDiscoveryProvidersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取所有发现Provider列表
func NewListDiscoveryProvidersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDiscoveryProvidersLogic {
	return &ListDiscoveryProvidersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDiscoveryProvidersLogic) ListDiscoveryProviders() (resp *types.ListProvidersResp, err error) {
	// 调用RPC服务获取完整的Provider列表（包括内置providers和数据库中的）
	rpcResp, err := l.svcCtx.IoRpc.ListDiscoveryProviders(l.ctx, &io.Empty{})
	if err != nil {
		logx.Errorw("调用RPC获取Provider列表失败", logx.Field("error", err.Error()))
		return &types.ListProvidersResp{Code: 1, Msg: "获取Provider列表失败"}, nil
	}

	// 转换RPC响应为API响应
	list := make([]types.ProviderMetadata, 0, len(rpcResp.Data))
	for _, p := range rpcResp.Data {
		metadata := types.ProviderMetadata{
			Id:       p.Id,
			Name:     p.Name,
			Category: p.Category,
		}
		if p.Description != nil {
			metadata.Description = *p.Description
		}
		if p.Version != nil {
			metadata.Version = *p.Version
		}
		if p.Icon != nil {
			metadata.Icon = *p.Icon
		}
		list = append(list, metadata)
	}

	return &types.ListProvidersResp{
		Code: 0,
		Msg:  "success",
		Data: list,
	}, nil
}
