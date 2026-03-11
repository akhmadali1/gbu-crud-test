// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

// User -.
type User struct {
	ID        int64      `json:"id"`
	UserName  string     `json:"user_name" validate:"required,max=255"`
	FirstName string     `json:"first_name" validate:"required,max=255"`
	LastName  string     `json:"last_name" validate:"max=255"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UpdateUser -.
type UpdateUser struct {
	ID        int64      `json:"id" validate:"required"`
	UserName  string     `json:"user_name" validate:"required,max=255"`
	FirstName string     `json:"first_name" validate:"required,max=255"`
	LastName  string     `json:"last_name" validate:"max=255"`
}

// DeleteUserDTO -.
type DeleteUserDTO struct {
	ID int64 `json:"id" validate:"required"`
}
