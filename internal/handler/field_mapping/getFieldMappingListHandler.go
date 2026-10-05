package field_mapping

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/field_mapping"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFieldMappingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFieldMappingListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := field_mapping.NewGetFieldMappingListLogic(r.Context(), svcCtx)
		resp, err := l.GetFieldMappingList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
