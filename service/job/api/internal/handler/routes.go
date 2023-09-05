// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"job/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/job/serchjob",
				Handler: jobsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/job/tagadd",
				Handler: tagHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/job/getcompanysum",
				Handler: companyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/job/getjob",
				Handler: jobHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}