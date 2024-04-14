// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
	"net"
	"time"
)

type Deck struct {
	ID        string
	UserID    string
	Title     string
	Cards     []byte
	Subject   sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	ID          string
	UserID      string
	Fingerprint []byte
	IP          net.IP
	ExpiresAt   time.Time
	CreatedAt   time.Time
}

type User struct {
	ID          string
	Username    string
	Email       string
	Password    []byte
	Name        string
	Description sql.NullString
	Picture     sql.NullString
	CreatedAt   time.Time
}
