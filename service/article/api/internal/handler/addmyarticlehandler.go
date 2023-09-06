package handler

import (
	"net/http"

	"article/api/internal/logic"
	"article/api/internal/svc"
	"article/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func addmyarticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddArtRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddmyarticleLogic(r.Context(), svcCtx)
		resp, err := l.Addmyarticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
