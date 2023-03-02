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

//func GetTransactionsById(ctx context.Context, page int) (*[]Domain, int, int, int64, error) {
//	return nil, 0, 0, 0, nil
//}
//
//func UpdateTransactionById(ctx context.Context, data *Domain, id string, updatedBy string) (*Domain, error) {
//	return nil, nil
//}
//
//func DeleteTransactionById(ctx context.Context, id string) (string, error) {
//	return "", nil
//}
