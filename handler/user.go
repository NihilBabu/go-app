package handler

import (
	"encoding/json"
	"github.com/NihilBabu/micro/dto"
	"github.com/NihilBabu/micro/model"
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/google/uuid"
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

		json.NewEncoder(w).Encode("error due to " + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func (h *Handlers) addUser(w http.ResponseWriter, r *http.Request)  {

	decoder := json.NewDecoder(r.Body)
	var userDto dto.UserDTO
	err := decoder.Decode(&userDto)
	if err != nil {
		json.NewEncoder(w).Encode("pase error")
		return
	}

	user := model.User{
		Email: userDto.Email,
		Name: userDto.Username,
		Password: userDto.Password,
		Model: model.Model{
			ID: uuid.New().String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	user, err = h.db.SaveUser(user)

	if err !=nil{
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}