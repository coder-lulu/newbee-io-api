package config

import (
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type ListConfigRequest struct {
	Page        *uint64 `json:"page,optional"`
	PageSize    *uint64 `json:"pageSize,optional"`
	ServiceName *string `json:"serviceName,optional"`
	Category    *string `json:"category,optional"`
	ConfigGroup *string `json:"configGroup,optional"`
	Keyword     *string `json:"keyword,optional"`
}

func ListConfigHandler(s *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ListConfigRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		data, err := s.IoRpc.ListConfig(r.Context(), &ioclient.ListConfigReq{Page: req.Page, PageSize: req.PageSize, ServiceName: req.ServiceName, Category: req.Category, ConfigGroup: req.ConfigGroup, Keyword: req.Keyword})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		items := make([]map[string]any, 0, len(data.Data))
		for _, item := range data.Data {
			items = append(items, mapConfigItem(item))
		}
		payload := map[string]any{"data": items}
		payload["total"] = data.GetTotal()
		httpx.OkJsonCtx(r.Context(), w, map[string]any{"code": 0, "msg": "success", "data": payload})
	}
}

type ListAuditLogRequest struct {
	Page       *uint64 `json:"page,optional"`
	PageSize   *uint64 `json:"pageSize,optional"`
	ConfigKey  *string `json:"configKey,optional"`
	ChangeType *string `json:"changeType,optional"`
	StartTime  *int64  `json:"startTime,optional"`
	EndTime    *int64  `json:"endTime,optional"`
}

func ListAuditLogHandler(s *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ListAuditLogRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		data, err := s.IoRpc.ListAuditLog(r.Context(), &ioclient.ListAuditLogReq{Page: req.Page, PageSize: req.PageSize, ConfigKey: req.ConfigKey, ChangeType: req.ChangeType, StartTime: req.StartTime, EndTime: req.EndTime})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		items := make([]map[string]any, 0, len(data.Data))
		for _, item := range data.Data {
			items = append(items, mapConfigAuditLog(item))
		}
		payload := map[string]any{"data": items}
		payload["total"] = data.GetTotal()
		httpx.OkJsonCtx(r.Context(), w, map[string]any{"code": 0, "msg": "success", "data": payload})
	}
}

type GetConfigHistoryRequest struct {
	ConfigKey *string `json:"configKey,optional"`
}

func GetConfigHistoryHandler(s *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GetConfigHistoryRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		data, err := s.IoRpc.GetConfigHistory(r.Context(), &ioclient.GetConfigHistoryReq{ConfigKey: req.ConfigKey})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		items := make([]map[string]any, 0, len(data.Data))
		for _, item := range data.Data {
			items = append(items, mapConfigAuditLog(item))
		}
		payload := map[string]any{"data": items}
		httpx.OkJsonCtx(r.Context(), w, map[string]any{"code": 0, "msg": "success", "data": payload})
	}
}

type GetConfigRequest struct {
	ConfigKey *string `json:"configKey,optional"`
}

func GetConfigHandler(s *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GetConfigRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		data, err := s.IoRpc.GetConfig(r.Context(), &ioclient.GetConfigReq{ConfigKey: req.ConfigKey})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		payload := map[string]any{"data": mapConfigItem(data.Data)}
		httpx.OkJsonCtx(r.Context(), w, map[string]any{"code": 0, "msg": "success", "data": payload})
	}
}
func mapConfigItem(v *ioclient.ConfigItem) map[string]any {
	if v == nil {
		return nil
	}
	return map[string]any{"id": v.GetId(), "tenantId": v.GetTenantId(), "configKey": v.GetConfigKey(), "configValue": v.GetConfigValue(), "valueType": v.GetValueType(), "category": v.GetCategory(), "serviceName": v.GetServiceName(), "description": v.GetDescription(), "defaultValue": v.GetDefaultValue(), "version": v.GetVersion(), "status": v.GetStatus(), "isReadonly": v.GetIsReadonly(), "isSensitive": v.GetIsSensitive(), "scope": v.GetScope(), "configGroup": v.GetConfigGroup(), "createdAt": v.GetCreatedAt(), "updatedAt": v.GetUpdatedAt()}
}
func mapConfigAuditLog(v *ioclient.ConfigAuditLog) map[string]any {
	if v == nil {
		return nil
	}
	return map[string]any{"id": v.GetId(), "tenantId": v.GetTenantId(), "configKey": v.GetConfigKey(), "oldValue": v.GetOldValue(), "newValue": v.GetNewValue(), "changeType": v.GetChangeType(), "changedBy": v.GetChangedBy(), "changedByName": v.GetChangedByName(), "serviceName": v.GetServiceName(), "category": v.GetCategory(), "configGroup": v.GetConfigGroup(), "changeReason": v.GetChangeReason(), "ipAddress": v.GetIpAddress(), "userAgent": v.GetUserAgent(), "oldVersion": v.GetOldVersion(), "newVersion": v.GetNewVersion(), "isRollback": v.GetIsRollback(), "rollbackFromLogId": v.GetRollbackFromLogId(), "createdAt": v.GetCreatedAt()}
}
