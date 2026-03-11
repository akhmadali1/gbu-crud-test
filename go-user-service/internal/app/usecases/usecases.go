package usecases

import (
	"go-user-service/config"
	persistentUser "go-user-service/internal/user/repo/persistent"
	"go-user-service/internal/user/usecase/user"
	"go-user-service/pkg/postgres"
)

// UseCases holds all the use cases for DI

type UseCases struct {
	User *user.UseCase
	// Add other services here
}

// New -.
func New(pg *postgres.Postgres, cfg *config.Config) *UseCases {
	return &UseCases{
		User: user.New(
			persistentUser.New(pg),
			cfg,
		),
		// Add other services here
	}
}
