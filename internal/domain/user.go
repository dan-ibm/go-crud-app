package domain

import "time"

type User struct {
	ID        int       `json:"-" db:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
