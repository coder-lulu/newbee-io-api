package output_task

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/output_task"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateOutputTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateOutputTaskReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := output_task.NewUpdateOutputTaskLogic(r.Context(), svcCtx)
		resp, err := l.UpdateOutputTask(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
