package response

import (
	"time"
)

type Transaction struct {
	ID        uint      `json:"id"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	Amount    int64     `json:"amount"`
	Notes     string    `json:"notes"`
	Date      time.Time `json:"date"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
