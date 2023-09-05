package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xiangxiang/jackaroo/service/company/api/api/internal/logic"
	"xiangxiang/jackaroo/service/company/api/api/internal/svc"
	"xiangxiang/jackaroo/service/company/api/api/internal/types"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
