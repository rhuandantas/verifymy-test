package server

import (
	"context"
	"fmt"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rhuandantas/verifymy-test/internal/config"
	"github.com/rhuandantas/verifymy-test/internal/log"
	"github.com/rhuandantas/verifymy-test/internal/server/handlers"
	"go.uber.org/zap"
)

type HttpServer struct {
	appName     *string
	host        string
	Server      *echo.Echo
	config      config.ConfigProvider
	logger      log.SimpleLogger
	userHandler *handlers.UserHandler
}

// NewAPIServer creates the main server with all configurations necessary
func NewAPIServer(config config.ConfigProvider, logger log.SimpleLogger, userHandler *handlers.UserHandler) *HttpServer {
	appName := config.GetStringOrDefault("app.name", "verify-my-service")
	host := config.GetStringOrDefault("server.host", "0.0.0.0:8080")

	app := echo.New()

	app.HideBanner = true
	app.HidePort = true

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info(
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	return &HttpServer{
		appName:     &appName,
		host:        host,
		Server:      app,
		config:      config,
		logger:      logger,
		userHandler: userHandler,
	}
}

func (hs *HttpServer) RegisterHandlers() {
	hs.userHandler.RegisterRoutes(hs.Server)
}

// Start starts an application on specific port
func (hs *HttpServer) Start() {
	ctx := context.Background()
	hs.logger.Info(ctx, fmt.Sprintf("Starting a server at http://%s", hs.host))
	err := hs.Server.Start(hs.host)
	if err != nil {
		hs.logger.Error(ctx, errorx.Decorate(err, "Failed to start the server"))
		return
	}
}
