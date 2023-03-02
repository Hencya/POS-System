package request

type Transaction struct {
	Amount int64  `json:"amount"`
	Notes  string `json:"notes"`
	Type   string `json:"type"`
}
