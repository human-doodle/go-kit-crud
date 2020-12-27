package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
)

//ErrRepo is ...
var ErrRepo = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *mgo.Database
	logger log.Logger
}

//NewRepo is ...
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
	}
	fmt.Println("User Created:", user.Email, user.Password, user.City, user.Age)

	return nil
}

/*
func (repo *repo) GetUser(ctx context.Context, id string) (profile Profile, errr error) {
	coll := repo.db.C("bloguser")
	data := []Profile{}
	d := Profile{}
	err := coll.Find(bson.M{"userid": id}).Select(bson.M{}).All(&data)
	if err != nil {
		fmt.Println("Error occured in GetuserById")
		return d, err
	}
	//fmt.Println(data)
	return data[0], nil
}*/
func (repo *repo) GetUser(ctx context.Context) (interface{}, error) {
	coll := repo.db.C("bloguser")
	data := []User{}
	err := coll.Find(bson.M{}).Select(bson.M{"email": 1, "city": 1, "age": 1}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return data, nil
}


/*

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
*/

func (repo *repo) UpdateUser(ctx context.Context, id int, user User) error {
	coll := repo.db.C("bloguser")
	err := coll.Update(bson.M{"userid": id}, bson.M{"$set": bson.M{"email": user.Email}})
	if err != nil {
		fmt.Println("Error occured inside update user repo")
		return err
	} else {
		return nil
	}

}