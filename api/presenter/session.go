package presenter

import (
	"time"
)

type LogInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	ID        string    `json:"id"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type AuthResponse struct {
	Session  Session `json:"session"`
	Username string  `json:"username"`
}
