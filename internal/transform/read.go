package transform

import (
	"encoding/json"
	"math"

	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"
)

func InputTask(v *ioclient.InputTaskInfo) types.InputTaskInfo {
	if v == nil {
		return types.InputTaskInfo{}
	}
	return types.InputTaskInfo{BaseIDInfo: types.BaseIDInfo{Id: v.Id, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt},
		Name:             v.GetTaskName(),
		Description:      v.GetDescription(),
		TaskType:         v.GetInputSource(),
		TaskStatus:       v.GetTaskStatus(),
		Priority:         int(v.GetPriority()),
		TimeoutSeconds:   int(v.GetTimeoutSeconds()),
		TaskConfig:       v.GetTaskConfig(),
		ValidationConfig: v.GetValidationConfig(),
		OutputTargets:    v.GetOutputTargets(),
		ScheduledAt:      v.ScheduledAt,
		StartedAt:        v.StartedAt,
		CompletedAt:      v.CompletedAt,
		TotalRecords:     v.GetTotalRecords(),
		ProcessedRecords: v.GetProcessedRecords(),
		SuccessRecords:   v.GetSuccessRecords(),
		FailedRecords:    v.GetFailedRecords(),
		ErrorMessage:     v.ErrorMessage,
		ErrorDetails:     v.GetErrorDetails(),
		MaxRetry:         int(v.GetMaxRetry()),
		RetryCount:       int(v.GetRetryCount()),
		NextRetryAt:      v.NextRetryAt,
		DiscoveryPoolId:  v.DiscoveryPoolId,
		WorkerId:         v.WorkerId,
		ProgressPercent:  taskProgress(v.GetProcessedRecords(), v.GetTotalRecords()),
		ProgressMessage:  v.GetProgressMessage(),
		Metadata:         v.GetMetadata()}
}

func OutputTask(v *ioclient.OutputTaskInfo) types.OutputTaskInfo {
	if v == nil {
		return types.OutputTaskInfo{}
	}
	return types.OutputTaskInfo{BaseIDInfo: types.BaseIDInfo{Id: v.Id, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt},
		Name:             v.GetTaskName(),
		OutputType:       v.GetOutputTarget(),
		ProgressPercent:  taskProgress(v.GetProcessedRecords(), v.GetTotalRecords()),
		TaskStatus:       v.GetTaskStatus(),
		PushTargets:      v.GetTargetConfig(),
		ScheduledAt:      v.ScheduledAt,
		StartedAt:        v.StartedAt,
		CompletedAt:      v.CompletedAt,
		TotalRecords:     v.GetTotalRecords(),
		ProcessedRecords: v.GetProcessedRecords(),
		SuccessRecords:   v.GetSuccessRecords(),
		FailedRecords:    v.GetFailedRecords(),
		ErrorMessage:     v.ErrorMessage,
		Metadata:         v.GetMetadata()}
}

func FieldMapping(v *ioclient.FieldMappingInfo) types.FieldMappingInfo {
	if v == nil {
		return types.FieldMappingInfo{}
	}
	return types.FieldMappingInfo{BaseIDInfo: types.BaseIDInfo{Id: v.Id, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt},
		MappingName:         v.GetMappingName(),
		Description:         v.GetDescription(),
		MappingType:         v.GetMappingType(),
		IsActive:            v.GetIsActive(),
		SourceField:         v.GetSourceField(),
		SourceFieldPath:     v.GetSourceFieldPath(),
		SourceDataType:      v.GetSourceDataType(),
		SourceFormat:        v.GetSourceFormat(),
		TargetField:         v.GetTargetField(),
		TargetFieldPath:     v.GetTargetFieldPath(),
		TargetDataType:      v.GetTargetDataType(),
		TargetFormat:        v.GetTargetFormat(),
		TransformType:       v.GetTransformType(),
		TransformConfig:     v.GetTransformConfig(),
		DefaultValue:        v.GetDefaultValue(),
		AllowNull:           v.GetAllowNull(),
		IsRequired:          v.GetIsRequired(),
		ValidationRules:     v.GetValidationRules(),
		ValidationRegex:     v.GetValidationRegex(),
		LookupTable:         v.GetLookupTable(),
		LookupCaseSensitive: v.GetLookupCaseSensitive(),
		ConditionRules:      v.GetConditionRules(),
		Priority:            int(v.GetPriority()),
		SortOrder:           int(v.GetSortOrder()),
		DiscoveryPoolId:     v.DiscoveryPoolId,
		InputTaskId:         v.InputTaskId,
		OutputTaskId:        v.OutputTaskId,
		UsageCount:          v.GetUsageCount(),
		SuccessCount:        v.GetSuccessCount(),
		FailedCount:         v.GetFailedCount(),
		LastUsedAt:          v.LastUsedAt,
		LastError:           v.LastError,
		LastErrorAt:         v.LastErrorAt,
		Metadata:            v.GetMetadata()}
}

func DataTarget(v *ioclient.DataTargetInfo) types.DataTargetInfo {
	if v == nil {
		return types.DataTargetInfo{}
	}
	return types.DataTargetInfo{Id: v.GetId(),
		CreatedAt:        v.GetCreatedAt(),
		UpdatedAt:        v.GetUpdatedAt(),
		Status:           uint8(v.GetStatus()),
		TargetName:       v.GetTargetName(),
		TargetType:       v.GetTargetType(),
		TargetSchema:     v.TargetSchema,
		ValidationRules:  v.ValidationRules,
		UniqueKeys:       v.UniqueKeys,
		ConnectionConfig: v.ConnectionConfig,
		Description:      v.Description,
		Metadata:         v.Metadata}
}

func DiscoveryPool(v *ioclient.DiscoveryPoolInfo) types.DiscoveryPoolInfo {
	if v == nil {
		return types.DiscoveryPoolInfo{}
	}
	return types.DiscoveryPoolInfo{BaseIDInfo: types.BaseIDInfo{Id: v.Id, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt},
		Name:            v.GetName(),
		Description:     v.GetDescription(),
		DiscoveryType:   v.GetDiscoveryType(),
		PoolStatus:      v.GetPoolStatus(),
		DiscoveryConfig: v.GetDiscoveryConfig(),
		Schedule:        v.GetSchedule(),
		BatchSize:       int(v.GetBatchSize()),
		ConcurrentLimit: int(v.GetConcurrentLimit()),
		MaxRetry:        int(v.GetMaxRetry()),
		RetryInterval:   int(v.GetRetryInterval()),
		FieldMapping:    v.GetFieldMapping(),
		TotalRuns:       v.GetTotalRuns(),
		SuccessRuns:     v.GetSuccessRuns(),
		FailedRuns:      v.GetFailedRuns(),
		LastRunAt:       v.LastRunAt,
		LastSuccessAt:   v.LastSuccessAt,
		LastError:       v.LastError,
		ApprovalStatus:  v.GetApprovalStatus(),
		ApprovedBy:      v.ApprovedBy,
		ApprovedAt:      v.ApprovedAt,
		RejectionReason: v.RejectionReason,
		Metadata:        v.GetMetadata()}
}

func DiscoveryTemplate(v *ioclient.DiscoveryTemplateInfo) types.DiscoveryTemplateInfo {
	if v == nil {
		return types.DiscoveryTemplateInfo{}
	}
	return types.DiscoveryTemplateInfo{Id: v.GetId(),
		CreatedAt:             v.GetCreatedAt(),
		UpdatedAt:             v.GetUpdatedAt(),
		Status:                uint8(v.GetStatus()),
		TenantId:              v.GetTenantId(),
		TemplateName:          v.GetTemplateName(),
		TemplateCode:          v.GetTemplateCode(),
		Description:           v.Description,
		Version:               v.GetVersion(),
		TemplateType:          v.GetTemplateType(),
		DiscoveryConfig:       v.DiscoveryConfig,
		FieldMappingTemplates: v.FieldMappingTemplates,
		ValidationRules:       v.ValidationRules,
		IsPublic:              v.GetIsPublic(),
		IsSystem:              v.GetIsSystem(),
		UsageCount:            v.GetUsageCount(),
		Tags:                  v.Tags,
		Metadata:              v.Metadata}
}

func TaskLog(v *ioclient.TaskLogInfo) types.TaskLogInfo {
	if v == nil {
		return types.TaskLogInfo{}
	}
	return types.TaskLogInfo{BaseIDInfo: types.BaseIDInfo{Id: v.Id, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt},
		TaskId:   v.GetTaskId(),
		TaskType: v.GetTaskType(),
		Level:    v.GetLogLevel(),
		Message:  v.GetLogMessage(),
		Details:  v.GetLogDetail()}
}

func MappingLog(v *ioclient.MappingLogInfo) types.MappingLogInfo {
	if v == nil {
		return types.MappingLogInfo{}
	}
	level := "info"
	if v.GetTransformStatus() == "failed" {
		level = "error"
	}
	details, _ := json.Marshal(map[string]string{"sourceValue": v.GetSourceValue(), "targetValue": v.GetTargetValue(), "transformStatus": v.GetTransformStatus()})
	message := v.GetErrorMessage()
	if message == "" {
		message = v.GetTransformStatus()
	}
	return types.MappingLogInfo{BaseIDInfo: types.BaseIDInfo{Id: v.Id, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt},
		FieldMappingId: v.GetFieldMappingId(),
		Level:          level,
		Message:        message,
		Details:        string(details)}
}

func WorkerMetrics(v *ioclient.WorkerMetricsInfo) types.WorkerMetricsInfo {
	if v == nil {
		return types.WorkerMetricsInfo{}
	}
	return types.WorkerMetricsInfo{Id: v.GetId(),
		CreatedAt:          v.GetCreatedAt(),
		UpdatedAt:          v.GetUpdatedAt(),
		WorkerId:           v.GetWorkerId(),
		CpuUsagePercent:    v.GetCpuUsage(),
		MemoryUsagePercent: v.GetMemoryUsage(),
		CurrentTaskCount:   v.GetCurrentTasks(),
		MetricTime:         v.GetLastHeartbeat(),
		Metadata:           v.Metadata}
}

// taskProgress derives the displayed percentage from the persisted record counts.
func taskProgress(processed, total int64) int {
	if total <= 0 || processed <= 0 {
		return 0
	}
	if processed >= total {
		return 100
	}
	return int(math.Round(float64(processed) / float64(total) * 100))
}
