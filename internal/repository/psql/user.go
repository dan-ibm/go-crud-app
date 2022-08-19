package psql

import (
	"fmt"
	"github.com/dan-ibm/go-crud-app/internal/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password, created_at) values ($1, $2, $3, $4) RETURNING id", "users")

	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Auth) GetUser(email, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", "users")
	err := r.db.Get(&user, query, email, password)

	return user, err
}
