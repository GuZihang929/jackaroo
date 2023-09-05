package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"job/api/internal/logic"
	"job/api/internal/svc"
	"job/api/internal/types"
)

func jobsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JobsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJobsLogic(r.Context(), svcCtx)
		resp, err := l.Jobs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
