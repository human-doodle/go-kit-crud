package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//NewHTTPServer is ...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	))
    /*
	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		//decodeEmailReq,
		decodeGetUserRequest,
		encodeResponse,
	))*/

	r.Methods("GET").Path("/getusers").Handler(httptransport.NewServer(
		endpoints.GetUser,
		//decodeEmailReq,
		decodeGetUserRequest,
		encodeResponse,
	))

     /*
	r.Methods("PUT").Path("/update").Handler(httptransport.NewServer(
		endpoints.UpdateUser,
		decodeUserReq,
		encodeResponse,
	))*/

	r.Methods("PUT").Path("/update").Handler(httptransport.NewServer(
		endpoints.UpdateUser,
		decodeUpdateUserRequest,
		encodeResponse,
	))


	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
