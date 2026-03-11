//nolint:godot,revive // Disable for swagger.
package httprouter

import (
	"go-user-service/config"
	"go-user-service/internal/app/router/http/middleware"
	"go-user-service/internal/app/usecases"
	userV1 "go-user-service/internal/user/controller/http/v1"
	"go-user-service/pkg/logger"
	"go-user-service/pkg/postgres"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, pg *postgres.Postgres, cfg *config.Config, l logger.Interface) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Usecases
	allUsecases := usecases.New(pg, cfg)

	// Routers
	apiV1Group := app.Group("/v1")
	{
		userV1.NewUserRoutes(apiV1Group, cfg, allUsecases.User, l)
	}
}
