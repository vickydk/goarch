package http

import (
	"goarch/pkg/interface/interceptor"
	"goarch/pkg/shared/config"
	"goarch/pkg/shared/logger"
	"goarch/pkg/shared/utils"
	"goarch/pkg/shared/utils/context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupMiddleware(server *echo.Echo, cfg *config.Config) {
	server.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Request().Header.Get(echo.HeaderXRequestID)
			if len(reqId) == 0 {
				reqId = utils.GenerateThreadId()
			}

			ctxSess := context.New(logger.GetLogger()).
				SetXRequestID(reqId).
				SetAppName("clerked.API").
				SetAppVersion("0.0").
				SetPort(cfg.Apps.HttpPort).
				SetSrcIP(c.RealIP()).
				SetURL(c.Request().URL.String()).
				SetMethod(c.Request().Method).
				SetHeader(c.Request().Header)

			ctxSess.Lv1("Incoming Request")

			c.Set(context.AppSession, ctxSess)

			return h(c)
		}
	})
	interceptor := interceptor.New()
	server.Use(interceptor.ValidateAccess())

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.HEAD, echo.OPTIONS},
		AllowHeaders: []string{
			"Content-Type", "Origin", "Accept", "Authorization", "Content-Length", "X-Requested-With",
			"OS-Type", "Device-Name", "Device-Serial", "OS-Version", "App-Version",
		},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}
