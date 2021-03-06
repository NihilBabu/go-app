package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"github.com/NihilBabu/micro/storage"
)

const message = "hlo welcome"

type Handlers struct {
	logger *log.Logger
	db     storage.Service
}

func New(logger *log.Logger, db storage.Service) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}

func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/user", h.benchmark(h.getUsers)).Methods("GET")
	mux.HandleFunc("/user", h.benchmark(h.addUser)).Methods("POST")
	mux.HandleFunc("/user/{userId}", h.benchmark(h.getUser)).Methods("GET")
}

func (h *Handlers) benchmark(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("Request processed for url %s %s in %s\n", r.Method ,r.URL.String(), time.Now().Sub(startTime))
		next(w, r)
	}
}
