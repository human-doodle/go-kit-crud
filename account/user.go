package account

import "context"

//User is ...
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
	Age      int    `json: "age"`
}

// Profile is ...
/*type Profile struct {
	Email string `json:"email"`
	City  string `json:"city"`
	Age   int    `json: "age"`
}*/

// Repository is ...
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	//GetUser(ctx context.Context, id string) (profile Profile, errr error)
	GetUser(ctx context.Context) (interface{}, error)
	UpdateUser(ctx context.Context, id int, user User) error
	//UpdateUser(ctx context.Context, user User) error
}
