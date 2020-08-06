package util

import (
	"github.com/NihilBabu/go-app/model"
	"github.com/NihilBabu/go-app/storage"
)

func InitialMigrate() {
	db := storage.GetDatabaseConnection()

	defer db.Close()
	db.AutoMigrate(&model.User{})
}
