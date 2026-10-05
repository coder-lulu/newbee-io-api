package config

import (
	"context"
	"encoding/json"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-rpc/ioclient"
	"google.golang.org/grpc"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeConfig struct {
	ioclient.Io
	t   *testing.T
	ctx context.Context
}

func (f *fakeConfig) ListConfig(ctx context.Context, r *ioclient.ListConfigReq, _ ...grpc.CallOption) (*ioclient.ListConfigResp, error) {
	if ctx != f.ctx || r.GetKeyword() != "demo.newbee." {
		f.t.Fatal("request/context lost")
	}
	total := uint64(1)
	key := "demo.newbee.test"
	return &ioclient.ListConfigResp{Total: &total, Data: []*ioclient.ConfigItem{{ConfigKey: &key}}}, nil
}
func TestListConfigEnvelope(t *testing.T) {
	r := httptest.NewRequest("POST", "/config/list", strings.NewReader(`{"page":1,"pageSize":20,"keyword":"demo.newbee."}`))
	r.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ListConfigHandler(&svc.ServiceContext{IoRpc: &fakeConfig{t: t, ctx: r.Context()}})(rec, r)
	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["code"] != float64(0) {
		t.Fatal(body)
	}
	data := body["data"].(map[string]any)
	if data["total"] != float64(1) || data["data"].([]any)[0].(map[string]any)["configKey"] != "demo.newbee.test" {
		t.Fatal(body)
	}
}
