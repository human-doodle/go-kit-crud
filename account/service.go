package account

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string, city string, age int) (string, error)
	GetUser(ctx context.Context, id string) (email string, city string, age int, errr error)
	UpdateUser(ctx context.Context, is string, email string, password string, city string, age int) (string, error)
}
