package presenter

import "time"

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"createdAt"`
}
