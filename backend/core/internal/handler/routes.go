// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	basic "github.com/pjimming/baize/core/internal/handler/basic"
	local "github.com/pjimming/baize/core/internal/handler/local"
	"github.com/pjimming/baize/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/baize/v1/local/module",
				Handler: local.GetModuleInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/baize/v1/local/packages",
				Handler: local.GetPackagesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/baize/v1/local/file",
				Handler: local.GetGoFilesHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/baize/v1/ping",
				Handler: basic.PingHandler(serverCtx),
			},
		},
	)
}