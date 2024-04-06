// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"time"
)

type Deck struct {
	ID        string
	Title     string
	Cards     []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	ID          string
	UserID      string
	Fingerprint []byte
	IP          []byte
	ExpiresAt   time.Time
	CreatedAt   time.Time
}

type User struct {
	ID        string
	Username  string
	Email     string
	Password  []byte
	CreatedAt time.Time
}
