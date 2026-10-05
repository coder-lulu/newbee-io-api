package discovery_pool

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/discovery_pool"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDiscoveryPoolListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDiscoveryPoolListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := discovery_pool.NewGetDiscoveryPoolListLogic(r.Context(), svcCtx)
		resp, err := l.GetDiscoveryPoolList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
