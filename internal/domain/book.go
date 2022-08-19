package domain

import (
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date" db:"publish_date"`
	Rating      int       `json:"rating"`
	UserId      int       `json:"user_id" db:"user_id"`
}

type BookInput struct {
	Title       *string    `json:"title"`
	PublishDate *time.Time `json:"publish_date"`
	Rating      *int       `json:"rating"`
}
