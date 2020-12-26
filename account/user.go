package account

import "context"

type User struct {
	ID       string `json:"id, omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
	Age      int    `json: "age"`
}

type Profile struct {
	Email string `json:"email"`
	City  string `json:"city"`
	Age   int    `json: "age"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (profile Profile, errr error)
	UpdateUser(ctx context.Context, user User) error
}
