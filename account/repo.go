package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB //Change for mongo
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {

	//change fr mongo

	// sql := `
	// 	INSERT INTO users (id, email, password)
	// 	VALUES ($1, $2, $3)`

	// if user.Email == "" || user.Password == "" {
	// 	return RepoErr
	// }

	// _, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string

	//change for mongo

	// err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return "", RepoErr
	}

	return email, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) error {

	//change fr mongo

	// update logic in mongo

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	// _, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}