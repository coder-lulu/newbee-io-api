package discovery_pool

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/discovery_pool"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApproveDiscoveryPoolHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApproveDiscoveryPoolReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := discovery_pool.NewApproveDiscoveryPoolLogic(r.Context(), svcCtx)
		resp, err := l.ApproveDiscoveryPool(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
