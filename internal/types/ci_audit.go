package types

type ListChangeHistoryReq struct {
	Page          int     `json:"page,optional"`
	PageSize      int     `json:"pageSize,optional"`
	CiID          *uint64 `json:"ciId,optional"`
	OperationType string  `json:"operationType,optional"`
	OperatorName  string  `json:"operatorName,optional"`
	Source        string  `json:"source,optional"`
	Status        string  `json:"status,optional"`
	NeedsApproval *bool   `json:"needsApproval,optional"`
}

type ListLifecycleStateReq struct {
	Page        int     `json:"page,optional"`
	PageSize    int     `json:"pageSize,optional"`
	CiID        *uint64 `json:"ciId,optional"`
	StateType   string  `json:"stateType,optional"`
	TriggerType string  `json:"triggerType,optional"`
	IsCurrent   *bool   `json:"isCurrent,optional"`
	IsTimeout   *bool   `json:"isTimeout,optional"`
	HasError    *bool   `json:"hasError,optional"`
}

type CiAuditListData struct {
	Data  []map[string]any `json:"data"`
	Total int              `json:"total"`
}

type CiAuditListResp struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data CiAuditListData `json:"data"`
}
