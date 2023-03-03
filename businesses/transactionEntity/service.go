package transactionEntity

import (
	"context"
	"fmt"
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

func (s *TransactionServices) GetTransactions(ctx context.Context, params ParamGetTransactions) (*[]Domain, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	if params.Page == 1 {
		params.Offset = 0
	} else {
		params.Offset = (params.Page - 1) * params.Limit
	}

	fmt.Println("params usecase ", params)
	res, totalData, err := s.TransactionRepository.GetTransactions(ctx, params)
	if err != nil {
		return &[]Domain{}, -1, -1, businesses.ErrNotFoundTransaction
	}

	return res, params.Offset, totalData, nil
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
