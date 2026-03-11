package persistent

import (
	"context"
	"fmt"
	"time"

	"go-user-service/internal/user/entity"
	"go-user-service/pkg/postgres"
)

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) GetAllData(ctx context.Context) ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("id, user_name, first_name, last_name, created_at, updated_at").
		From("users").
		Where("deleted_at IS NULL").
		OrderBy("id DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAllData - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAllData - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	datas := make([]entity.User, 0)
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID, &e.UserName, &e.FirstName, &e.LastName, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetAllData - rows.Scan: %w", err)
		}

		datas = append(datas, e)
	}

	return datas, nil
}

func (r *UserRepo) CreateData(ctx context.Context, req entity.User) (*entity.User, error) {
	sql, args, err := r.Builder.
		Insert("users").
		Columns(
			"user_name",
			"first_name",
			"last_name",
		).
		Values(
			req.UserName,
			req.FirstName,
			req.LastName,
		).
		Suffix("RETURNING user_name, first_name, last_name, created_at, updated_at, deleted_at").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("UserRepo - CreateData - ToSql: %w", err)
	}

	var data entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&data.UserName, &data.FirstName, &data.LastName, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - CreateData - QueryRow: %w", err)
	}

	return &data, nil
}

func (r *UserRepo) UpdateData(ctx context.Context, req entity.UpdateUser) (*entity.User, error) {
	sql, args, err := r.Builder.
		Update("users").
		Set("user_name", req.UserName).
		Set("first_name", req.FirstName).
		Set("last_name", req.LastName).
		Set("updated_at", time.Now()).
		Where("id = ?", req.ID).
		Suffix("RETURNING user_name, first_name, last_name, created_at, updated_at, deleted_at").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("UserRepo - UpdateData - ToSql: %w", err)
	}

	var data entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&data.UserName, &data.FirstName, &data.LastName, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - UpdateData - QueryRow: %w", err)
	}

	return &data, nil
}

func (r *UserRepo) DeleteData(ctx context.Context, id int64) (*entity.User, error) {
	sql, args, err := r.Builder.
		Update("users").
		Set("deleted_at", time.Now()).
		Where("id = ?", id).
		Suffix("RETURNING user_name, first_name, last_name, created_at, updated_at, deleted_at").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("UserRepo - DeleteData - ToSql: %w", err)
	}

	var data entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&data.UserName, &data.FirstName, &data.LastName, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - DeleteData - QueryRow: %w", err)
	}

	return &data, nil
}
