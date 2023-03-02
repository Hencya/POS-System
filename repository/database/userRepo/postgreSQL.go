package userRepo

import (
	"context"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"POS-System/businesses/userEntity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userEntity.Repository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (userEntity.Domain, error) {
	domain := userEntity.Domain{}
	rec := User{}
	err := r.db.Where("username = ?", username).First(&rec).Error
	if err != nil {
		return userEntity.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}
