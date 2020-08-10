package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handlers) getUsers(w http.ResponseWriter, r *http.Request) {
	//h.db.ExecContext(r.Context(), "select ")
	users, err := h.db.GetUsers()
	if err != nil {
		h.logger.Printf("data fetching failed due to %v\n", err)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *Handlers) getUser(w http.ResponseWriter, r *http.Request) {
	//h.db.ExecContext(r.Context(), "select ")
	vars := mux.Vars(r)
	users, err := h.db.GetUser(vars["userId"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("error due to" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
