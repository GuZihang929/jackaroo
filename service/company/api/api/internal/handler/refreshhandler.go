package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xiangxiang/jackaroo/service/company/api/api/internal/logic"
	"xiangxiang/jackaroo/service/company/api/api/internal/svc"
	"xiangxiang/jackaroo/service/company/api/api/internal/types"
)

func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshLogic(r.Context(), svcCtx)
		resp, err := l.Refresh(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
