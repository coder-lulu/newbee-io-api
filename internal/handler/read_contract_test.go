package handler_test

import (
	"context"
	"encoding/json"
	"github.com/coder-lulu/newbee-io-api/internal/handler/field_mapping"
	"github.com/coder-lulu/newbee-io-api/internal/handler/mapping_log"
	"github.com/coder-lulu/newbee-io-api/internal/handler/output_task"
	"github.com/coder-lulu/newbee-io-api/internal/handler/task_log"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
	"google.golang.org/grpc"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type readRPC struct {
	ioclient.Io
	t   *testing.T
	ctx context.Context
}

func (f *readRPC) GetOutputTaskList(ctx context.Context, r *ioclient.OutputTaskListReq, _ ...grpc.CallOption) (*ioclient.OutputTaskListResp, error) {
	if ctx != f.ctx || r.Page != 1 || r.PageSize != 20 {
		f.t.Fatalf("lost context/pagination: %d/%d", r.Page, r.PageSize)
	}
	data := make([]*ioclient.OutputTaskInfo, 20)
	for i := range data {
		id := uint64(i + 1)
		data[i] = &ioclient.OutputTaskInfo{Id: &id}
	}
	return &ioclient.OutputTaskListResp{Total: 30, Data: data}, nil
}
func (f *readRPC) GetFieldMappingList(ctx context.Context, r *ioclient.FieldMappingListReq, _ ...grpc.CallOption) (*ioclient.FieldMappingListResp, error) {
	if ctx != f.ctx || r.Page != 1 || r.PageSize != 20 {
		f.t.Fatalf("lost context/pagination: %d/%d", r.Page, r.PageSize)
	}
	data := make([]*ioclient.FieldMappingInfo, 20)
	for i := range data {
		id := uint64(i + 1)
		data[i] = &ioclient.FieldMappingInfo{Id: &id}
	}
	return &ioclient.FieldMappingListResp{Total: 30, Data: data}, nil
}
func (f *readRPC) GetTaskLogList(ctx context.Context, r *ioclient.TaskLogListReq, _ ...grpc.CallOption) (*ioclient.TaskLogListResp, error) {
	if ctx != f.ctx || r.Page != 1 || r.PageSize != 20 {
		f.t.Fatalf("lost context/pagination: %d/%d", r.Page, r.PageSize)
	}
	data := make([]*ioclient.TaskLogInfo, 20)
	for i := range data {
		id := uint64(i + 1)
		data[i] = &ioclient.TaskLogInfo{Id: &id}
	}
	return &ioclient.TaskLogListResp{Total: 30, Data: data}, nil
}
func (f *readRPC) GetMappingLogList(ctx context.Context, r *ioclient.MappingLogListReq, _ ...grpc.CallOption) (*ioclient.MappingLogListResp, error) {
	if ctx != f.ctx || r.Page != 1 || r.PageSize != 20 {
		f.t.Fatalf("lost context/pagination: %d/%d", r.Page, r.PageSize)
	}
	data := make([]*ioclient.MappingLogInfo, 20)
	for i := range data {
		id := uint64(i + 1)
		data[i] = &ioclient.MappingLogInfo{Id: &id}
	}
	return &ioclient.MappingLogListResp{Total: 30, Data: data}, nil
}
func TestGETReadListsWithoutOptionalFilters(t *testing.T) {
	for _, tc := range []struct {
		name    string
		handler func(*svc.ServiceContext) http.HandlerFunc
	}{{"output_task", output_task.GetOutputTaskListHandler}, {"field_mapping", field_mapping.GetFieldMappingListHandler}, {"task_log", task_log.GetTaskLogListHandler}, {"mapping_log", mapping_log.GetMappingLogListHandler}} {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/"+tc.name+"/list?page=1&pageSize=20", nil)
			w := httptest.NewRecorder()
			tc.handler(&svc.ServiceContext{IoRpc: &readRPC{t: t, ctx: r.Context()}})(w, r)
			var response struct {
				Code int `json:"code"`
				Data struct {
					Total uint64            `json:"total"`
					Data  []json.RawMessage `json:"data"`
				} `json:"data"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
			}
			if w.Code != 200 || response.Code != 0 || response.Data.Total != 30 || len(response.Data.Data) != 20 {
				t.Fatalf("failed GET contract: %s", w.Body.String())
			}
		})
	}
}
func TestPOSTPageQueryPreservesBodyFilters(t *testing.T) {
	r := httptest.NewRequest("POST", "/input_task/list?page=2&pageSize=20", strings.NewReader(`{"page":2,"pageSize":20,"name":"demo"}`))
	r.Header.Set("Content-Type", "application/json")
	var req types.GetInputTaskListReq
	if err := httpx.Parse(r, &req, true); err != nil {
		t.Fatal(err)
	}
	if req.Page != 2 || req.PageSize != 20 || req.Name == nil || *req.Name != "demo" {
		t.Fatalf("bad request: %+v", req)
	}
}
func TestGETOptionalFalseFilterIsPreserved(t *testing.T) {
	r := httptest.NewRequest("GET", "/field_mapping/list?page=1&pageSize=20&isActive=false", nil)
	var req types.GetFieldMappingListReq
	if err := httpx.Parse(r, &req, true); err != nil {
		t.Fatal(err)
	}
	if req.IsActive == nil || *req.IsActive {
		t.Fatal("false filter lost")
	}
}

func TestGETLogDetailPathID(t *testing.T) {
	r := pathvar.WithVars(httptest.NewRequest("GET", "/task_log/42", nil), map[string]string{"id": "42"})
	var req types.IDPathReq
	if err := httpx.Parse(r, &req, true); err != nil {
		t.Fatal(err)
	}
	if req.Id != 42 {
		t.Fatal("path ID lost")
	}
}
