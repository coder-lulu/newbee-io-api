package logic_test

import (
	"context"
	"errors"
	"github.com/coder-lulu/newbee-io-api/internal/logic/data_target"
	"github.com/coder-lulu/newbee-io-api/internal/logic/discovery_pool"
	"github.com/coder-lulu/newbee-io-api/internal/logic/discovery_template"
	"github.com/coder-lulu/newbee-io-api/internal/logic/field_mapping"
	"github.com/coder-lulu/newbee-io-api/internal/logic/input_task"
	"github.com/coder-lulu/newbee-io-api/internal/logic/mapping_log"
	"github.com/coder-lulu/newbee-io-api/internal/logic/output_task"
	"github.com/coder-lulu/newbee-io-api/internal/logic/task_log"
	"github.com/coder-lulu/newbee-io-api/internal/logic/worker_metrics"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"
	"google.golang.org/grpc"
	"testing"
)

type ctxKey struct{}
type fakeIO struct {
	ioclient.Io
	want context.Context
	t    *testing.T
	fail error
}

func (f *fakeIO) check(ctx context.Context) {
	f.t.Helper()
	if ctx != f.want || ctx.Value(ctxKey{}) != "tenant-permission-scope" {
		f.t.Fatal("original tenant/permission context lost")
	}
}
func (f *fakeIO) GetInputTaskList(ctx context.Context, req *ioclient.InputTaskListReq, _ ...grpc.CallOption) (*ioclient.InputTaskListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.InputTaskListResp{}, nil
}
func TestGetInputTaskList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := input_task.NewGetInputTaskListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetInputTaskList(&types.GetInputTaskListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetInputTaskList(&types.GetInputTaskListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetInputTaskById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.InputTaskInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.InputTaskInfo{}, nil
}
func TestGetInputTask(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := input_task.NewGetInputTaskLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetInputTask(&types.IDPathReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetInputTask(&types.IDPathReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetOutputTaskList(ctx context.Context, req *ioclient.OutputTaskListReq, _ ...grpc.CallOption) (*ioclient.OutputTaskListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.OutputTaskListResp{}, nil
}
func TestGetOutputTaskList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := output_task.NewGetOutputTaskListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetOutputTaskList(&types.GetOutputTaskListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetOutputTaskList(&types.GetOutputTaskListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetOutputTaskById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.OutputTaskInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.OutputTaskInfo{}, nil
}
func TestGetOutputTask(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := output_task.NewGetOutputTaskLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetOutputTask(&types.IDPathReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetOutputTask(&types.IDPathReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetFieldMappingList(ctx context.Context, req *ioclient.FieldMappingListReq, _ ...grpc.CallOption) (*ioclient.FieldMappingListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.FieldMappingListResp{}, nil
}
func TestGetFieldMappingList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := field_mapping.NewGetFieldMappingListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetFieldMappingList(&types.GetFieldMappingListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetFieldMappingList(&types.GetFieldMappingListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetFieldMappingById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.FieldMappingInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.FieldMappingInfo{}, nil
}
func TestGetFieldMapping(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := field_mapping.NewGetFieldMappingLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetFieldMapping(&types.IDPathReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetFieldMapping(&types.IDPathReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetDataTargetById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.DataTargetInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.DataTargetInfo{}, nil
}
func TestGetDataTargetById(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := data_target.NewGetDataTargetByIdLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetDataTargetById(&types.IDReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetDataTargetById(&types.IDReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetDataTargetList(ctx context.Context, req *ioclient.DataTargetListReq, _ ...grpc.CallOption) (*ioclient.DataTargetListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.DataTargetListResp{}, nil
}
func TestGetDataTargetList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := data_target.NewGetDataTargetListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetDataTargetList(&types.DataTargetListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetDataTargetList(&types.DataTargetListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetDiscoveryTemplateById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.DiscoveryTemplateInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.DiscoveryTemplateInfo{}, nil
}
func TestGetDiscoveryTemplateById(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := discovery_template.NewGetDiscoveryTemplateByIdLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetDiscoveryTemplateById(&types.IDReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetDiscoveryTemplateById(&types.IDReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetDiscoveryTemplateList(ctx context.Context, req *ioclient.DiscoveryTemplateListReq, _ ...grpc.CallOption) (*ioclient.DiscoveryTemplateListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.DiscoveryTemplateListResp{}, nil
}
func TestGetDiscoveryTemplateList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := discovery_template.NewGetDiscoveryTemplateListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetDiscoveryTemplateList(&types.DiscoveryTemplateListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetDiscoveryTemplateList(&types.DiscoveryTemplateListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetTaskLogList(ctx context.Context, req *ioclient.TaskLogListReq, _ ...grpc.CallOption) (*ioclient.TaskLogListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.TaskLogListResp{}, nil
}
func TestGetTaskLogList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := task_log.NewGetTaskLogListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetTaskLogList(&types.GetTaskLogListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetTaskLogList(&types.GetTaskLogListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetTaskLogById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.TaskLogInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.TaskLogInfo{}, nil
}
func TestGetTaskLog(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := task_log.NewGetTaskLogLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetTaskLog(&types.IDPathReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetTaskLog(&types.IDPathReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetMappingLogList(ctx context.Context, req *ioclient.MappingLogListReq, _ ...grpc.CallOption) (*ioclient.MappingLogListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.MappingLogListResp{}, nil
}
func TestGetMappingLogList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := mapping_log.NewGetMappingLogListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetMappingLogList(&types.GetMappingLogListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetMappingLogList(&types.GetMappingLogListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetMappingLogById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.MappingLogInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.MappingLogInfo{}, nil
}
func TestGetMappingLog(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := mapping_log.NewGetMappingLogLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetMappingLog(&types.IDPathReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetMappingLog(&types.IDPathReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetWorkerMetricsById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.WorkerMetricsInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.WorkerMetricsInfo{}, nil
}
func TestGetWorkerMetricsById(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := worker_metrics.NewGetWorkerMetricsByIdLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetWorkerMetricsById(&types.IDReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetWorkerMetricsById(&types.IDReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetWorkerMetricsList(ctx context.Context, req *ioclient.WorkerMetricsListReq, _ ...grpc.CallOption) (*ioclient.WorkerMetricsListResp, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.WorkerMetricsListResp{}, nil
}
func TestGetWorkerMetricsList(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := worker_metrics.NewGetWorkerMetricsListLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetWorkerMetricsList(&types.WorkerMetricsListReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetWorkerMetricsList(&types.WorkerMetricsListReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func (f *fakeIO) GetDiscoveryPoolById(ctx context.Context, req *ioclient.IDReq, _ ...grpc.CallOption) (*ioclient.DiscoveryPoolInfo, error) {
	f.check(ctx)
	if f.fail != nil {
		return nil, f.fail
	}
	return &ioclient.DiscoveryPoolInfo{}, nil
}
func TestGetDiscoveryPool(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := discovery_pool.NewGetDiscoveryPoolLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetDiscoveryPool(&types.IDPathReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetDiscoveryPool(&types.IDPathReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
func TestGetDiscoveryPoolStats(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "tenant-permission-scope")
	f := &fakeIO{want: ctx, t: t}
	l := discovery_pool.NewGetDiscoveryPoolStatsLogic(ctx, &svc.ServiceContext{IoRpc: f})
	response, err := l.GetDiscoveryPoolStats(&types.IDPathReq{})
	if err != nil || response == nil {
		t.Fatalf("empty RPC response must return data: %v", err)
	}
	f.fail = errors.New("tenant access denied")
	response, err = l.GetDiscoveryPoolStats(&types.IDPathReq{})
	if !errors.Is(err, f.fail) || response != nil {
		t.Fatal("RPC access error swallowed")
	}
}
