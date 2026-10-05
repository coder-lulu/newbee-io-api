package discoveryprovider

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/discoveryprovider"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 测试Provider连接
func TestProviderConnectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestConnectionReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := discoveryprovider.NewTestProviderConnectionLogic(r.Context(), svcCtx)
		resp, err := l.TestProviderConnection(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
