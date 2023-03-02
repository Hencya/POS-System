package userEntity

import (
	"context"
	"time"
)

type Domain struct {
	ID        uint
	Username  string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Login(ctx context.Context, username string, password string) (string, error)
}

type Repository interface {
	// Databases postgresql
	GetByUsername(ctx context.Context, username string) (Domain, error)
}
