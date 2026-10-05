package discoveryprovider

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/discoveryprovider"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取所有发现Provider列表
func ListDiscoveryProvidersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := discoveryprovider.NewListDiscoveryProvidersLogic(r.Context(), svcCtx)
		resp, err := l.ListDiscoveryProviders()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
