package v1

import (
	"go-user-service/internal/user/usecase"
	"go-user-service/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	t usecase.User
	l logger.Interface
	v *validator.Validate
}
