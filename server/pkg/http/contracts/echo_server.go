package contracts

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/pkg/http/config"
)

type EchoHttpServer interface {
	RunHttpServer(configEcho ...func(echo *echo.Echo)) error
	GracefulShutdown(ctx context.Context) error
	GetEchoInstance() *echo.Echo
	RouteBuilder() *RouteBuilder
	AddMiddlewares(middlewares ...echo.MiddlewareFunc)
	Cfg() *config.EchoHttpOptions
	SetupDefaultMiddlewares()
}
