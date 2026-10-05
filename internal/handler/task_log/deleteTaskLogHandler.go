package task_log

import (
	"net/http"

	"github.com/coder-lulu/newbee-io-api/internal/logic/task_log"
	"github.com/coder-lulu/newbee-io-api/internal/svc"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteTaskLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task_log.NewDeleteTaskLogLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTaskLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
