package v1

import (
	"go-user-service/config"
	"go-user-service/internal/user/usecase/user"
	"go-user-service/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewUserRoutes -.
func NewUserRoutes(apiV1Group fiber.Router, cfg *config.Config, t *user.UseCase, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}
	userGroup := apiV1Group.Group("/user")
	{
		userGroup.Get("", func(c *fiber.Ctx) error {
			return r.getAllData(c)
		})
		userGroup.Post("", func(c *fiber.Ctx) error {
			return r.createData(c)
		})
		userGroup.Put("", func(c *fiber.Ctx) error {
			return r.updateData(c)
		})
		userGroup.Delete("", func(c *fiber.Ctx) error {
			return r.deleteData(c)
		})
	}
}
