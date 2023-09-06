package handler

import (
	"net/http"

	"article/api/internal/logic"
	"article/api/internal/svc"
	"article/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func articleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArtRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), svcCtx)
		resp, err := l.Article(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
