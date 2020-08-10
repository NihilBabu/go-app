package storage

import (
	"log"

	"github.com/NihilBabu/micro/model"
)

type storage struct {
	logger *log.Logger
}

type Service interface {
	LoadTables()
	GetUsers() ([]model.User, error)
	SaveUser(model.User) (*model.User, error)
	GetUser(string) (model.User, error)
	Close() error
}
