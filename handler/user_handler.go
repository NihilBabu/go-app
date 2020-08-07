package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/NihilBabu/go-app/model"
	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {

	users, err := storageService.GetUsers()

	if err != nil {
		json.NewEncoder(w).Encode("ERROR")
	}
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	email := vars["email"]
	user := model.User{
		Name:     name,
		Email:    email,
		Password: name,
		Model: model.Model{
			CreatedAt: time.Now(),
			IsActive:  true,
			IsDeleted: false,
			ID:        uuid.New().String()},
	}

	user1, err := storageService.SaveUser(user)
	if err != nil {
		json.NewEncoder(w).Encode("User Saved failed" + err.Error())
	}
	json.NewEncoder(w).Encode(user1)

}
