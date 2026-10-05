package data_target

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/data_target"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDataTargetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DataTargetListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := data_target.NewGetDataTargetListLogic(r.Context(), svcCtx)
		resp, err := l.GetDataTargetList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]any{"code": 0, "msg": "success", "data": resp})
		}
	}
}
