// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"backstage/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/backstage/get24p",
				Handler: get24pHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/backstage/getseajob",
				Handler: getseajobHandler(serverCtx),
			},
		},
		//rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
