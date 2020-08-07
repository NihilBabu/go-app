package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NihilBabu/go-app/storage"
	"github.com/gorilla/mux"
)

var storageService storage.Service

// New returns an http handler for the url shortener.
func New(svc storage.Service) *mux.Router {
	r := mux.NewRouter()
	storageService = svc
	// &service{storage}

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user/{name}/{email}", createUser).Methods("POST")
	return r
}

type service struct {
	storage.Service
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("hiii")
}
