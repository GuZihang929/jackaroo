package handler

import (
	"net/http"

	"article/api/internal/logic"
	"article/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func myarticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMyarticleLogic(r.Context(), svcCtx)
		resp, err := l.Myarticle()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
