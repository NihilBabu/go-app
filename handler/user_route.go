package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NihilBabu/go-app/model"
	"github.com/NihilBabu/go-app/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetTtest(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("Hello")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := storage.GetDatabaseConnection()
	defer db.Close()

	var users []model.User
	db.Where("is_deleted =?", false).Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	email := vars["email"]
	user := &model.User{
		Model: model.Model{
			ID:        uuid.New().String(),
			IsActive:  true,
			IsDeleted: false,
		},
		Name:     name,
		Password: name,
		Email:    email,
	}
	db := storage.GetDatabaseConnection()
	defer db.Close()

	db.Create(user)
	json.NewEncoder(w).Encode(user)
}
