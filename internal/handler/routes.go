// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "github.com/iot-synergy/synergy-member-api/internal/handler/base"
	device "github.com/iot-synergy/synergy-member-api/internal/handler/device"
	"github.com/iot-synergy/synergy-member-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/queryLocalDevices",
					Handler: device.QueryLocalDevicesHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
