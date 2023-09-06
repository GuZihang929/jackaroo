// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"article/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/article/articles",
				Handler: articlesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/article/article",
				Handler: articleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/article/addmyarticle",
				Handler: addmyarticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/article/myarticle",
				Handler: myarticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/article/updatearticle",
				Handler: updatearticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/article/delarticle",
				Handler: delarticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/article/art",
				Handler: artHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/article/likeadd",
				Handler: likeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
