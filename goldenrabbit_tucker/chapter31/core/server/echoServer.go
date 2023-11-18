package core_server

import (
	core_router "gt/chap31/ex/core/router"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func EchoServerStart(address string) error {
	echoServer := echo.New()
	echoServer.Use(getCorsMiddleware())
	routerLoad(echoServer)
	return echoServer.Start(address)
}

func getCorsMiddleware() echo.MiddlewareFunc {
	return echoMiddleware.CORSWithConfig(getCorsConfig())
}

func getCorsConfig() echoMiddleware.CORSConfig {
	return echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderContentType},
		AllowCredentials: true,
		MaxAge:           60 * 10,
	}
}

func routerLoad(e *echo.Echo) {
	core_router.Load(e)
}
