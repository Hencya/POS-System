package userEntity

import (
	"context"
	"strings"
	"time"

	"POS-System/app/middlewares/auth"
	"POS-System/businesses"
	"POS-System/helpers"
)

type UserServices struct {
	UserRepository Repository
	JwtAuth        *auth.ConfigJWT
	ContextTimeout time.Duration
}

func NewUserServices(repoUser Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &UserServices{
		UserRepository: repoUser,
		JwtAuth:        auth,
		ContextTimeout: timeout,
	}
}

func (s *UserServices) Login(ctx context.Context, username string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrUsernamePasswordNotFound
	}

	doctorDomain, err := s.UserRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", businesses.ErrUsernameNotRegistered
	}

	if !helpers.ValidateHash(password, doctorDomain.Password) {
		return "", businesses.ErrPassword
	}

	token := s.JwtAuth.GenerateToken(doctorDomain.Username)

	return token, nil
}
