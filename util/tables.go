package util

import (
	"github.com/NihilBabu/go-app/storage"
)

func InitialMigrate(svc storage.Service) {

	svc.LoadTables()
	// db.AutoMigrate(&model.User{})
}
