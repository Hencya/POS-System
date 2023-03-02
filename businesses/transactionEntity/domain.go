package transactionEntity

import (
	"context"
	"time"
)

type Domain struct {
	ID        uint
	CreatedBy string
	UpdatedBy string
	Amount    int64
	Notes     string
	Date      time.Time
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateNewTransaction(ctx context.Context, data *Domain) (*Domain, error)
	//GetTransactionsById(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
	//UpdateTransactionById(ctx context.Context, data *Domain, id string, updatedBy string) (*Domain, error)
	//DeleteTransactionById(ctx context.Context, id string) (string, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewTransaction(ctx context.Context, data *Domain) (*Domain, error)
	//UpdateTransactionById(ctx context.Context, id string, data *Domain) (*Domain, error)
	//DeleteTransactionById(ctx context.Context, id string) (string, error)
	//GetTransactionsById(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
}
