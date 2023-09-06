package handler

import (
	"net/http"

	"backstage/api/internal/logic"
	"backstage/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getseajobHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetseajobLogic(r.Context(), svcCtx)
		resp, err := l.Getseajob()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
