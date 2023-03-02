package transactionRepo

import (
	"context"

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

//func (r *TransactionRepository) UpdateTransactionById(ctx context.Context, id string, data *transactionEntity.Domain) (*transactionEntity.Domain, error) {
//	return nil, nil
//}
//
//func (r *TransactionRepository) DeleteTransactionById(ctx context.Context, id string) (string, error) {
//	return "", nil
//}
//
//func (r *TransactionRepository) GetTransactionsById(ctx context.Context, offset, limit int) (*[]transactionEntity.Domain, int64, error) {
//	return nil, 0, nil
//}
