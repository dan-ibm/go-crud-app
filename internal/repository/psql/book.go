package psql

import (
	"fmt"
	"github.com/dan-ibm/go-crud-app/internal/domain"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Book struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *Book {
	return &Book{db}
}

func (r *Book) Create(userId int, book domain.BookInput) error {
	query := "INSERT INTO books (title, publish_date, rating, user_id) VALUES ($1, $2, $3, $4) RETURNING id"
	_, err := r.db.Exec(query, book.Title, book.PublishDate, book.Rating, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Book) GetById(id int) (domain.Book, error) {
	var book domain.Book

	query := "SELECT id, title, publish_date, rating, user_id FROM books WHERE id=$1"
	err := r.db.Get(&book, query, id)

	return book, err
}

func (r *Book) GetAll() ([]domain.Book, error) {
	var books []domain.Book

	query := "SELECT b.id, b.title, b.publish_date, b.rating, b.user_id FROM books b"
	err := r.db.Select(&books, query)

	return books, err
}

func (r *Book) GetAllByUser(userId int) ([]domain.Book, error) {
	var books []domain.Book

	query := "SELECT id, title, publish_date, rating, user_id FROM books WHERE user_id = $1"
	err := r.db.Select(&books, query, userId)

	return books, err
}

func (r *Book) Delete(id int) error {
	query := "DELETE FROM books WHERE id=$1"
	_, err := r.db.Exec(query, id)

	return err
}

func (r *Book) Update(id int, input domain.BookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.PublishDate != nil {
		setValues = append(setValues, fmt.Sprintf("publish_date=$%d", argId))
		args = append(args, *input.PublishDate)
		argId++
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *input.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE books SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
