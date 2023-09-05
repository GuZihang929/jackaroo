package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"job/api/internal/logic"
	"job/api/internal/svc"
)

func companyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCompanyLogic(r.Context(), svcCtx)
		resp, err := l.Company()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
