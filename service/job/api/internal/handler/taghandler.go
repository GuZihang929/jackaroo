package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"job/api/internal/logic"
	"job/api/internal/svc"
	"job/api/internal/types"
)

func tagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TagAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTagLogic(r.Context(), svcCtx)
		resp, err := l.Tag(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
