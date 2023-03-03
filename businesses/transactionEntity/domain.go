package transactionEntity

import (
	"context"
	"time"
)

type Domain struct {
	ID        uint
	CreatedBy string
	UpdatedBy string
	Amount    float64
	Notes     string
	Date      time.Time
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ParamGetTransactions struct {
	Page       int
	Offset     int
	Limit      int
	Amount     int
	TypeAmount string
	Type       string
	Sort       string
}

type Service interface {
	CreateNewTransaction(ctx context.Context, data *Domain) (*Domain, error)
	GetTransactions(ctx context.Context, params ParamGetTransactions) (*[]Domain, int, int64, error)
	GetTransactionById(ctx context.Context, id uint) (*Domain, error)
	UpdateTransactionById(ctx context.Context, data *Domain, id uint) (*Domain, error)
	DeleteTransactionById(ctx context.Context, id uint) (string, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewTransaction(ctx context.Context, data *Domain) (*Domain, error)
	GetTransactions(ctx context.Context, params ParamGetTransactions) (*[]Domain, int64, error)
	GetTransactionById(ctx context.Context, id uint) (*Domain, error)
	UpdateTransactionById(ctx context.Context, id uint, data *Domain) (*Domain, error)
	DeleteTransactionById(ctx context.Context, id uint) (string, error)
}
