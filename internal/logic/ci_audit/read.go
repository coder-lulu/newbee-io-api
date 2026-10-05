package ciaudit

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/coder-lulu/newbee-common/v2/middleware/keys"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ent"
	"github.com/coder-lulu/newbee-io-rpc/ent/cichangehistory"
	"github.com/coder-lulu/newbee-io-rpc/ent/cilifecyclestate"
)

// Require the authenticated middleware's canonical tenant, never a request-body tenant.
func tenantID(ctx context.Context) (uint64, error) {
	id, err := strconv.ParseUint(keys.NewContextManager().GetTenantID(ctx), 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("valid tenant context required")
	}
	return id, nil
}

func pagination(page, size int) (int, int) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	if size > 100 {
		size = 100
	}
	if page > 1000000 {
		page = 1000000
	}
	return (page - 1) * size, size
}

func ListChangeHistory(ctx context.Context, db *ent.Client, req *types.ListChangeHistoryReq) (*types.CiAuditListResp, error) {
	tenant, err := tenantID(ctx)
	if err != nil {
		return nil, err
	}
	q := db.CiChangeHistory.Query().Where(cichangehistory.TenantIDEQ(tenant))
	if req.CiID != nil {
		q.Where(cichangehistory.CiIDEQ(*req.CiID))
	}
	if req.OperationType != "" {
		q.Where(cichangehistory.OperationTypeEQ(req.OperationType))
	}
	if req.OperatorName != "" {
		q.Where(cichangehistory.OperatorNameContains(req.OperatorName))
	}
	if req.Source != "" {
		q.Where(cichangehistory.SourceEQ(req.Source))
	}
	if req.Status != "" {
		q.Where(cichangehistory.StatusEQ(req.Status))
	}
	if req.NeedsApproval != nil {
		q.Where(cichangehistory.NeedsApprovalEQ(*req.NeedsApproval))
	}
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	offset, limit := pagination(req.Page, req.PageSize)
	rows, err := q.Order(ent.Desc(cichangehistory.FieldCreatedAt), ent.Desc(cichangehistory.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	data := make([]map[string]any, 0, len(rows))
	for _, row := range rows {
		data = append(data, mapChangeHistory(row))
	}
	return &types.CiAuditListResp{Code: 0, Msg: "success", Data: types.CiAuditListData{Data: data, Total: total}}, nil
}

func ListLifecycleState(ctx context.Context, db *ent.Client, req *types.ListLifecycleStateReq) (*types.CiAuditListResp, error) {
	tenant, err := tenantID(ctx)
	if err != nil {
		return nil, err
	}
	q := db.CiLifecycleState.Query().Where(cilifecyclestate.TenantIDEQ(tenant))
	if req.CiID != nil {
		q.Where(cilifecyclestate.CiIDEQ(*req.CiID))
	}
	if req.StateType != "" {
		q.Where(cilifecyclestate.StateTypeEQ(req.StateType))
	}
	if req.TriggerType != "" {
		q.Where(cilifecyclestate.TriggerTypeEQ(req.TriggerType))
	}
	if req.IsCurrent != nil {
		q.Where(cilifecyclestate.IsCurrentEQ(*req.IsCurrent))
	}
	if req.IsTimeout != nil {
		q.Where(cilifecyclestate.IsTimeoutEQ(*req.IsTimeout))
	}
	if req.HasError != nil {
		q.Where(cilifecyclestate.HasErrorEQ(*req.HasError))
	}
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	offset, limit := pagination(req.Page, req.PageSize)
	rows, err := q.Order(ent.Desc(cilifecyclestate.FieldEnteredAt), ent.Desc(cilifecyclestate.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	data := make([]map[string]any, 0, len(rows))
	for _, row := range rows {
		data = append(data, mapLifecycleState(row))
	}
	return &types.CiAuditListResp{Code: 0, Msg: "success", Data: types.CiAuditListData{Data: data, Total: total}}, nil
}

func mapChangeHistory(v *ent.CiChangeHistory) map[string]any {
	return map[string]any{"id": v.ID, "operationId": v.OperationID, "ciId": v.CiID, "ciTypeId": v.CiTypeID, "operationType": v.OperationType, "operationName": v.OperationName, "operatorId": v.OperatorID, "operatorName": v.OperatorName, "source": v.Source, "sourceDetail": v.SourceDetail, "changeReason": v.ChangeReason, "status": v.Status, "affectedCount": v.AffectedCount, "needsApproval": v.NeedsApproval, "isApproved": v.IsApproved, "canRollback": v.CanRollback, "createdAt": v.CreatedAt.Format(time.RFC3339)}
}

func mapLifecycleState(v *ent.CiLifecycleState) map[string]any {
	exitedAt := ""
	if !v.ExitedAt.IsZero() {
		exitedAt = v.ExitedAt.Format(time.RFC3339)
	}
	return map[string]any{"id": v.ID, "stateId": v.StateID, "ciId": v.CiID, "ciTypeId": v.CiTypeID, "stateName": v.StateName, "stateType": v.StateType, "previousState": v.PreviousState, "enteredAt": v.EnteredAt.Format(time.RFC3339), "exitedAt": exitedAt, "durationSeconds": v.DurationSeconds, "triggerType": v.TriggerType, "triggeredBy": v.TriggeredBy, "triggeredByName": v.TriggeredByName, "isCurrent": v.IsCurrent, "isFinal": v.IsFinal, "isTimeout": v.IsTimeout, "hasError": v.HasError, "errorMessage": v.ErrorMessage, "canRetry": v.CanRetry}
}
