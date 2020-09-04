package sqlite

import (
	"github.com/NihilBabu/micro/model"
	"github.com/NihilBabu/micro/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
type Sqlite struct{ *gorm.DB }

func New(path string) (storage.Service, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &Sqlite{db}, nil
}

func (db *Sqlite) Close() error { return db.Close() }


func (s Sqlite) LoadTables() {
	s.AutoMigrate(model.User{})
}

func (s Sqlite) GetUsers() ([]model.User, error) {
	var user []model.User
	err := s.Where("is_deleted =?",false).Find(&user).Error
	return user, err
}

func (s Sqlite) SaveUser(user model.User) (model.User, error) {
	err := s.Create(user).Error
	return user, err
}

func (s Sqlite) GetUser(userId string) (model.User, error) {
	var user model.User
	err := s.Where("id =? and is_deleted =?", userId,false).Find(&user).Error
	return user, err
}

func (s Sqlite) DeleteUser(userId string) (error)  {
	err := s.Find("id =? and is_deleted =?", userId,false).UpdateColumn("is_deleted", true).Error

	return err
}