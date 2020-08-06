package main

import (
	"github.com/NihilBabu/go-app/handler"
	"github.com/gorilla/mux"
)

func GetAllRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.GetTtest).Methods("GET")
	r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{name}/{email}", handler.CreateUser).Methods("POST")
	return r
}
