package request

import "time"

type Transaction struct {
	Amount int64     `json:"amount"`
	Notes  string    `json:"notes"`
	Type   string    `json:"type"`
	Date   time.Time `json:"date"`
}
