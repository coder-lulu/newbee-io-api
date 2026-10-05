package transform

import (
	"github.com/coder-lulu/newbee-io-rpc/ioclient"
	"testing"
)

func ptr[T any](v T) *T { return &v }
func TestReadNamesAndOptionalFields(t *testing.T) {
	item := InputTask(&ioclient.InputTaskInfo{TaskName: ptr("demo-input"), InputSource: ptr("file"), TaskType: ptr("manual")})
	if item.Name != "demo-input" || item.TaskType != "file" || item.StartedAt != nil {
		t.Fatal(item)
	}
	log := TaskLog(&ioclient.TaskLogInfo{LogMessage: ptr("mapped"), LogLevel: ptr("warn")})
	if log.Message != "mapped" || log.Level != "warn" {
		t.Fatal(log)
	}
	worker := WorkerMetrics(&ioclient.WorkerMetricsInfo{WorkerId: ptr("worker-beijing-01"), CpuUsage: ptr(12.5)})
	if worker.WorkerId != "worker-beijing-01" || worker.CpuUsagePercent != 12.5 {
		t.Fatal(worker)
	}
}
func TestMappingFailurePreservesDetails(t *testing.T) {
	log := MappingLog(&ioclient.MappingLogInfo{TransformStatus: ptr("failed"), ErrorMessage: ptr("invalid field"), SourceValue: ptr("bad")})
	if log.Level != "error" || log.Message != "invalid field" || log.Details == "" {
		t.Fatal(log)
	}
}

func TestTaskProgressUsesRecordCounts(t *testing.T) {
	for _, tc := range []struct {
		name             string
		processed, total int64
		want             int
	}{{"completed", 30, 30, 100}, {"partial", 3, 10, 30}, {"unknown total", 10, 0, 0}, {"empty", 0, 0, 0}, {"over total", 20, 10, 100}, {"negative", -1, 10, 0}} {
		t.Run(tc.name, func(t *testing.T) {
			input := InputTask(&ioclient.InputTaskInfo{ProcessedRecords: ptr(tc.processed), TotalRecords: ptr(tc.total)})
			output := OutputTask(&ioclient.OutputTaskInfo{ProcessedRecords: ptr(tc.processed), TotalRecords: ptr(tc.total)})
			if input.ProgressPercent != tc.want || output.ProgressPercent != tc.want || input.TotalRecords != tc.total || output.TotalRecords != tc.total {
				t.Fatalf("input=%+v output=%+v", input, output)
			}
		})
	}
}
