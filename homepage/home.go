package homepage

import (
	"encoding/json"
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

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	//h.db.ExecContext(r.Context(), "select ")
	users, err := h.db.GetUsers()
	if err != nil {
		h.logger.Printf("data fetching failed due to %v\n", err)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte(users))
	json.NewEncoder(w).Encode(users)
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("Request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)

	}
}

func New(logger *log.Logger, db storage.Service) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}
