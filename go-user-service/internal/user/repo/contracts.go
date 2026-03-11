// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"
	"go-user-service/internal/user/entity"
)

type (
	// UserRepo -.
	UserRepo interface {
		GetAllData(ctx context.Context) ([]entity.User, error)
		CreateData(ctx context.Context, req entity.User) (*entity.User, error)
		UpdateData(ctx context.Context, req entity.UpdateUser) (*entity.User, error)
		DeleteData(ctx context.Context, id int64) (*entity.User, error)
	}
)
