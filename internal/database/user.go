package database

import (
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/autobrr/autobrr/internal/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) FindByUsername(username string) (*domain.User, error) {
	query := `SELECT username, password FROM users WHERE username = ?`

	row := r.db.QueryRow(query, username)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var user domain.User

	if err := row.Scan(&user.Username, &user.Password); err != nil {
		log.Error().Err(err).Msg("could not scan user to struct")
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Store(user domain.User) error {
	query := `INSERT INTO users (username, password) VALUES (?, ?)`

	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error executing query")
		return err
	}

	return nil
}
