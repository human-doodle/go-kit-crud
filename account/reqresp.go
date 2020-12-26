package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	//CreateUserRequest is ...
	CreateUserRequest struct {
		Email    string `json :"email"`
		Password string `json : "password"`
		City     string `json:"city"`
		Age      int    `json: "age"`
	}

	//CreateUserResponse is ...
	CreateUserResponse struct {
		Ok string `json : "ok"`
	}

	//GetUserRequest is...
	GetUserRequest struct {
		ID string `json : "id"`
	}

	//GetUserResponse is ...
	GetUserResponse struct {
		Email string `json : "email"`
		City  string `json:"city"`
		Age   int    `json: "age"`
	}

	//UpdateUserRequest is ...
	UpdateUserRequest struct {
		ID       string `json : "id"`
		Email    string `json :"email"`
		Password string `json : "password"`
		City     string `json:"city"`
		Age      int    `json: "age"`
	}

	//UpdateUserResponse is...
	UpdateUserResponse struct {
		Ok string `json : "ok"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		ID: vars["id"],
	}
	return req, nil
}
