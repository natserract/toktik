package http

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/natserract/toktik/pkg/http/config"
	"github.com/natserract/toktik/pkg/http/contracts"
	customMiddleware "github.com/natserract/toktik/pkg/http/middlewares"
)

type echoHttpServer struct {
	echo         *echo.Echo
	config       *config.EchoHttpOptions
	routeBuilder *contracts.RouteBuilder
}

func NewEchoHttpServer(
	config *config.EchoHttpOptions,
) contracts.EchoHttpServer {
	e := echo.New()
	e.HideBanner = true

	return &echoHttpServer{
		echo:         e,
		config:       config,
		routeBuilder: contracts.NewRouteBuilder(e),
	}
}

func (s *echoHttpServer) RunHttpServer(
	configEcho ...func(echo *echo.Echo),
) error {
	s.echo.Server.ReadTimeout = 15 * time.Second
	s.echo.Server.WriteTimeout = 15 * time.Second
	s.echo.Server.MaxHeaderBytes = 1 << 20

	if len(configEcho) > 0 {
		echoFunc := configEcho[0]
		if echoFunc != nil {
			configEcho[0](s.echo)
		}
	}

	// https://echo.labstack.com/guide/http_server/
	return s.echo.Start(s.config.Port)
}

func (s *echoHttpServer) Cfg() *config.EchoHttpOptions {
	return s.config
}

func (s *echoHttpServer) RouteBuilder() *contracts.RouteBuilder {
	return s.routeBuilder
}

func (s *echoHttpServer) AddMiddlewares(middlewares ...echo.MiddlewareFunc) {
	if len(middlewares) > 0 {
		s.echo.Use(middlewares...)
	}
}

func (s *echoHttpServer) GracefulShutdown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *echoHttpServer) SetupDefaultMiddlewares() {
	s.echo.Use(middleware.Recover())
	// log errors and information
	s.echo.Use(customMiddleware.LogMiddleware())
	s.echo.Use(middleware.BodyLimit("2M"))
	s.echo.Use(middleware.RequestID())
}

func (s *echoHttpServer) GetEchoInstance() *echo.Echo {
	return s.echo
}
