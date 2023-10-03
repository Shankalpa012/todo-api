package serve

import (
	"context"
	"fmt"
	"todo/bootstraps"
	"todo/controller"
	"todo/infrastructure"
	"todo/lib"
	"todo/middleware"
	"todo/routes"
	"todo/service"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module exports for fx
var Module = fx.Options(
	infrastructure.Module,
	controller.Module,
	service.Module,
	routes.Module,
	bootstraps.Module,
	middleware.Module,
	lib.Module,
	fx.Invoke(HttpServer),
)

func HttpServer(lifecycle fx.Lifecycle, h *infrastructure.Handler, logger lib.Logger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info(`Starting application in :8080`)
				logger.Info(`+-----------------------+`)
				logger.Info(`|       TODO API      |`)
				logger.Info(`+-----------------------+`)
				logger := zap.Must(zap.NewProduction())
				defer logger.Sync()
				logger.Info("Started Form Main.go!")
				go h.Gin.Run("localhost:8080")
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Warn(`Stopping Application!!!!`)
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
