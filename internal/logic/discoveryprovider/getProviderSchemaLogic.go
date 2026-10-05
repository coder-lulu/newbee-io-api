package discoveryprovider

import (
	"context"

	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/types/io"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProviderSchemaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取指定Provider的Schema
func NewGetProviderSchemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProviderSchemaLogic {
	return &GetProviderSchemaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProviderSchemaLogic) GetProviderSchema(req *types.GetProviderSchemaReq) (resp *types.GetProviderSchemaResp, err error) {
	// 调用RPC服务获取Provider Schema
	rpcResp, err := l.svcCtx.IoRpc.GetProviderSchema(l.ctx, &io.GetProviderSchemaReq{
		ProviderId: req.ProviderId,
	})
	if err != nil {
		logx.Errorw("调用RPC获取Provider Schema失败", logx.Field("error", err.Error()), logx.Field("provider_id", req.ProviderId))
		return &types.GetProviderSchemaResp{Code: 1, Msg: "获取Provider Schema失败"}, nil
	}

	// 转换RPC响应为API响应
	if rpcResp.Data == nil {
		return &types.GetProviderSchemaResp{
			Code: 0,
			Msg:  "success",
			Data: nil,
		}, nil
	}

	schema := &types.ProviderSchemaInfo{
		Id:            rpcResp.Data.Id,
		ProviderId:    rpcResp.Data.ProviderId,
		ProviderName:  rpcResp.Data.ProviderName,
		Category:      rpcResp.Data.Category,
		Description:   safeString(rpcResp.Data.Description),
		Version:       safeString(rpcResp.Data.Version),
		IconUrl:       safeString(rpcResp.Data.IconUrl),
		IsActive:      rpcResp.Data.IsActive,
		IsBuiltin:     rpcResp.Data.IsBuiltin,
		ExecutionMode: rpcResp.Data.ExecutionMode,
	}

	// 转换参数定义
	schema.ParameterSchema = make([]types.ParameterDefinition, 0, len(rpcResp.Data.ParameterSchema))
	for _, p := range rpcResp.Data.ParameterSchema {
		param := types.ParameterDefinition{
			Name:        p.Name,
			Label:       p.Label,
			Type:        p.Type,
			Required:    p.Required,
			Placeholder: safeString(p.Placeholder),
			Description: safeString(p.Description),
		}
		if p.DefaultValue != nil {
			param.DefaultValue = *p.DefaultValue
		}
		if p.Options != nil {
			param.Options = p.Options
		}
		schema.ParameterSchema = append(schema.ParameterSchema, param)
	}

	// 转换字段定义
	schema.FieldSchema = make([]types.FieldDefinition, 0, len(rpcResp.Data.FieldSchema))
	for _, f := range rpcResp.Data.FieldSchema {
		field := types.FieldDefinition{
			Name:        f.Name,
			Label:       f.Label,
			DataType:    f.DataType,
			Required:    f.Required,
			Description: safeString(f.Description),
			Example:     safeString(f.Example),
		}
		schema.FieldSchema = append(schema.FieldSchema, field)
	}

	return &types.GetProviderSchemaResp{
		Code: 0,
		Msg:  "success",
		Data: schema,
	}, nil
}

// 辅助函数：安全地解引用字符串指针
func safeString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
