package router

import (
	"github.com/gorilla/mux"
	"go_boilerplate/app/controller"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	hello := &controller.HelloHandler{
		CustomeMessage: "Hello",
	}

	r.Handle("/hello", hello).Methods("GET")

	return r
}
