package user

import (
	"context"
	"fmt"

	"go-user-service/config"
	"go-user-service/internal/user/entity"
	userRepo "go-user-service/internal/user/repo"
	sharedDto "go-user-service/shared/dto"
)

// UseCase -.
type UseCase struct {
	repo userRepo.UserRepo
	cfg  *config.Config
}

// New -.
func New(r userRepo.UserRepo, cfg *config.Config) *UseCase {
	return &UseCase{
		repo: r,
		cfg:  cfg,
	}
}

func (uc *UseCase) GetAllData(ctx context.Context) ([]entity.User, *sharedDto.APIError) {
	datas, err := uc.repo.GetAllData(ctx)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - GetAllData - repo.GetAllData: %v", err),
		}
	}

	return datas, nil
}

func (uc *UseCase) CreateData(ctx context.Context, req entity.User) (*entity.User, *sharedDto.APIError) {
	data, err := uc.repo.CreateData(ctx, req)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - CreateData - repo.CreateData: %v", err),
		}
	}

	return data, nil
}

func (uc *UseCase) UpdateData(ctx context.Context, req entity.UpdateUser) (*entity.User, *sharedDto.APIError) {
	data, err := uc.repo.UpdateData(ctx, req)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - UpdateData - repo.UpdateData: %v", err),
		}
	}

	return data, nil
}

func (uc *UseCase) DeleteData(ctx context.Context, id int64) (*entity.User, *sharedDto.APIError) {
	data, err := uc.repo.DeleteData(ctx, id)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - DeleteData - repo.DeleteData: %v", err),
		}
	}

	return data, nil
}
