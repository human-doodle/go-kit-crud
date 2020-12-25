package account

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string, city string, age int) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
	UpdateUser(ctx context.Context, email string, password string, city string, age int) (string, error)
}
