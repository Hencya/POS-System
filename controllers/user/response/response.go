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
