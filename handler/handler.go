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
	mux.HandleFunc("/", h.benchmark(h.getUsers))
	mux.HandleFunc("/user/{userId}", h.benchmark(h.getUser))
}

func (h *Handlers) benchmark(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("Request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}
