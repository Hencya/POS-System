package response

import (
	"time"
)

type Users struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
