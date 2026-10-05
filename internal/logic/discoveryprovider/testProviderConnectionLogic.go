package discoveryprovider

import (
	"context"
	"encoding/json"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/types/io"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestProviderConnectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 测试Provider连接
func NewTestProviderConnectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestProviderConnectionLogic {
	return &TestProviderConnectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestProviderConnectionLogic) TestProviderConnection(req *types.TestConnectionReq) (resp *types.TestConnectionResp, err error) {
	// 将config序列化为JSON字符串
	configJSON, err := json.Marshal(req.Config)
	if err != nil {
		logx.Errorw("序列化配置失败", logx.Field("error", err.Error()))
		return &types.TestConnectionResp{
			Code: 1,
			Msg:  "配置格式错误",
			Data: &types.TestConnectionData{
				Success: false,
				Message: "配置格式错误: " + err.Error(),
			},
		}, nil
	}

	// 调用RPC服务测试连接
	rpcResp, err := l.svcCtx.IoRpc.TestProviderConnection(l.ctx, &io.TestConnectionReq{
		ProviderId: req.ProviderId,
		Config:     string(configJSON),
	})
	if err != nil {
		logx.Errorw("调用RPC测试Provider连接失败", logx.Field("error", err.Error()), logx.Field("provider_id", req.ProviderId))
		return &types.TestConnectionResp{
			Code: 1,
			Msg:  "测试连接失败",
			Data: &types.TestConnectionData{
				Success: false,
				Message: "测试连接失败: " + err.Error(),
			},
		}, nil
	}

	// 转换RPC响应为API响应
	message := ""
	if rpcResp.Message != nil {
		message = *rpcResp.Message
	}

	return &types.TestConnectionResp{
		Code: 0,
		Msg:  "success",
		Data: &types.TestConnectionData{
			Success: rpcResp.Success,
			Message: message,
		},
	}, nil
}
