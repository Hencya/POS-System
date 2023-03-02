package repository

import (
	"gorm.io/gorm"

	"POS-System/businesses/transactionEntity"
	"POS-System/businesses/userEntity"
	"POS-System/repository/database/transactionRepo"
	"POS-System/repository/database/userRepo"
)

func NewUserRepository(db *gorm.DB) userEntity.Repository {
	return userRepo.NewUserRepository(db)
}

func NewTransactionRepository(db *gorm.DB) transactionEntity.Repository {
	return transactionRepo.NewTransactionRepository(db)
}
