package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	// fmt.Println("create user mongo repo", db)/
	err := repo.db.C("bloguser").Insert(user)
	if err != nil {
		fmt.Println("Error occured inside CreateUser in repo")
		return err
	} else {
		fmt.Println("User Created:", user.Email, user.Password, user.City, user.Age)
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (interface{}, error) {
	coll := repo.db.C("bloguser")
	data := []User{}
	err := coll.Find(bson.M{"userid": id}).Select(bson.M{}).All(&data)
	if err != nil {
		fmt.Println("Error occured in GetuserById")
		return "", err
	}
	return email,nil
}


func (repo *repo) UpdateUser(ctx context.Context, user User) error {
	err := repo.db.C("bloguser").Update(
		bson.M{"id": user.ID},
		bson.D{
			{"$set", bson.D{{"password", user.Password}}},
			{"$set", bson.D{{"email", user.Email}}},
			{"$set", bson.D{{"city", user.City}}},
			{"$set", bson.D{{"age", user.Age}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
