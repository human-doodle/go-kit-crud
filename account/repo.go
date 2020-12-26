package account

import (
	"context"
	"fmt"
	"errors"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)


var RepoErr = errors.New("Unable to handle Repo Request")


type repo struct {
	db     *mgo.Database
	logger log.Logger
}

func NewRepo(db *mgo.Database, logger log.Logger) (Repository, error) {
	return &repo{
	db:     db,
	logger: log.With(logger, "repo", "mongodb"),
	}, nil
	}

func (repo *repo) CreateUser(ctx context.Context, user User) error {

		fmt.Println("create user mongo repo", db)
		err := db.C("bloguser").Insert(user)
		if err != nil {
		fmt.Println("Error occured inside CreateUser in repo")
		return err
		} else {
		fmt.Println("User Created:", user.Email,user.Password,user.City,user.Age)
		}
		return nil
		}
	


func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string

	//change for mongo

	// err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	/*if err != nil {
		return "", RepoErr
	}
*/
	return email, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) error {

	//change fr mongo

	// update logic in mongo

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	// _, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	/*if err != nil {
		return err
	}*/
	return nil
}