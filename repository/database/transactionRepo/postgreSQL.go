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

func (r *TransactionRepository) GetTransactions(ctx context.Context, offset, limit int) (*[]transactionEntity.Domain, int64, error) {
	var totalData int64
	domain := []transactionEntity.Domain{}
	rec := []Transaction{}

	r.db.Find(&rec).Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return nil, 0, err
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

	//if err := r.db.First(&rec, "id = ?", id).Updates(&recData).Error; err != nil {
	//	return &transactionEntity.Domain{}, err
	//}

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
