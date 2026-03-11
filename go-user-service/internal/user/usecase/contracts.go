// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"go-user-service/internal/user/entity"
	sharedDto "go-user-service/shared/dto"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// User -.
	User interface {
		GetAllData(ctx context.Context) ([]entity.User, *sharedDto.APIError)
		CreateData(ctx context.Context, req entity.User) (*entity.User, *sharedDto.APIError)
		UpdateData(ctx context.Context, req entity.UpdateUser) (*entity.User, *sharedDto.APIError)
		DeleteData(ctx context.Context, id int64) (*entity.User, *sharedDto.APIError)
	}
)
