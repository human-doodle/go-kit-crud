package account

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints is...
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	UpdateUser endpoint.Endpoint
}

//MakeEndpoints is...
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password, req.City, req.Age)
		return CreateUserResponse{Ok: ok}, err
	}
}

/*
func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, city, age, err := s.GetUser(ctx, req.ID)

		return GetUserResponse{
			Email: email,
			City:  city,
			Age:   age,
		}, err
	}
} */
func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		data, err := s.GetUser(ctx)
		return GetUserResponse{Data: data, Err: err}, nil
	}

}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		ok, err := s.UpdateUser(ctx, req.ID, req.Email, req.Password, req.City, req.Age)
		return UpdateUserResponse{Ok: ok}, err
	}
}

// func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(UpdateUserRequest)
// 		rc := req.user
// 		 i, err := strconv.Atoi(rc.ID)
// 		 err = s.UpdateUser(ctx, i, req.user)
// 		return UpdateUserResponse{Err: err}, nil
// 	}
// }
