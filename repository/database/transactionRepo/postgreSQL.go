package transactionRepo

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"POS-System/businesses/transactionEntity"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transactionEntity.Repository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) CreateNewTransaction(ctx context.Context, data *transactionEntity.Domain) (*transactionEntity.Domain, error) {
	domain := transactionEntity.Domain{}
	rec := Transaction{}
	copier.Copy(&rec, &data)
	err := r.db.Create(&rec).Error
	if err != nil {
		return nil, err
	}
	copier.Copy(&domain, &rec)
	return &domain, nil
}

func (r *TransactionRepository) GetTransactions(ctx context.Context, params transactionEntity.ParamGetTransactions) (*[]transactionEntity.Domain, int64, error) {
	var totalData int64
	var err error
	domain := []transactionEntity.Domain{}
	rec := []Transaction{}

	r.db.Find(&rec).Count(&totalData)
	if params.Type != "" && params.TypeAmount != "" {
		switch params.TypeAmount {
		case "max":
			err = r.db.Where("type = ? AND amount <= ?", params.Type, params.Amount).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		case "min":
			err = r.db.Where("type = ? AND amount >= ?", params.Type, params.Amount).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		}
	} else if params.Type == "" && params.TypeAmount != "" {
		switch params.TypeAmount {
		case "max":
			err = r.db.Where("amount <= ?", params.Amount).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		case "min":
			err = r.db.Where("amount >= ?", params.Amount).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		}
	} else if params.Type != "" && params.TypeAmount == "" {
		err = r.db.Where("type = ?", params.Type).Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
			Find(&rec).Error
	} else {
		err = r.db.Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
			Find(&rec).Error
	}

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query transactions: %w", err)
	}

	copier.Copy(&domain, &rec)

	return &domain, totalData, nil
}

func (r *TransactionRepository) UpdateTransactionById(ctx context.Context, id uint, data *transactionEntity.Domain) (*transactionEntity.Domain, error) {
	domain := transactionEntity.Domain{}
	rec := Transaction{}
	recData := Transaction{}

	copier.Copy(&recData, &data)

	if err := r.db.First(&rec, "id = ?", id).Updates(&recData).Error; err != nil {
		fmt.Println(err)
		return &transactionEntity.Domain{}, err
	}

	copier.Copy(&domain, &rec)
	return &domain, nil
}

func (r *TransactionRepository) GetTransactionById(ctx context.Context, id uint) (*transactionEntity.Domain, error) {
	domain := transactionEntity.Domain{}
	rec := Transaction{}

	if err := r.db.Where("id = ?", id).First(&rec).Error; err != nil {
		return &transactionEntity.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return &domain, nil
}

func (r *TransactionRepository) DeleteTransactionById(ctx context.Context, id uint) (string, error) {
	rec := Transaction{}

	if err := r.db.Where("id = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Transaction was Deleted", nil
}
