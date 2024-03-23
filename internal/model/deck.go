package model

import "time"

type Deck struct {
	ID        int64
	Title     string
	Cards     [][2]string
	CreatedAt time.Time
	UpdatedAt time.Time
}
