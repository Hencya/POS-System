package transactionEntity

import (
	"context"
	"time"

	"POS-System/app/middlewares/auth"
	"POS-System/businesses"
)

type TransactionServices struct {
	TransactionRepository Repository
	JwtAuth               *auth.ConfigJWT
	ContextTimeout        time.Duration
}

func NewTransactionServices(repoTransaction Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &TransactionServices{
		TransactionRepository: repoTransaction,
		JwtAuth:               auth,
		ContextTimeout:        timeout,
	}
}

func (s *TransactionServices) CreateNewTransaction(ctx context.Context, data *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	res, err := s.TransactionRepository.CreateNewTransaction(ctx, data)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	return res, nil
}

func (s *TransactionServices) GetTransactions(ctx context.Context, page int) (*[]Domain, int, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	var offset int
	limit := 5
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	res, totalData, err := s.TransactionRepository.GetTransactions(ctx, offset, limit)
	if err != nil {
		return &[]Domain{}, -1, -1, -1, businesses.ErrNotFoundTransaction
	}

	return res, offset, limit, totalData, nil
}

func (s *TransactionServices) UpdateTransactionById(ctx context.Context, data *Domain, id uint) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()
	
	dataUpdated, err := s.TransactionRepository.UpdateTransactionById(ctx, id, data)
	if err != nil {
		return &Domain{}, businesses.ErrInternalServer
	}
	return dataUpdated, nil
}

func (s *TransactionServices) DeleteTransactionById(ctx context.Context, id uint) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	res, err := s.TransactionRepository.DeleteTransactionById(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundTransaction
	}
	return res, nil
}

func (s *TransactionServices) GetTransactionById(ctx context.Context, id uint) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	result, err := s.TransactionRepository.GetTransactionById(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}
