package mysql

import (
	"log"

	"github.com/NihilBabu/micro/model"
	"github.com/NihilBabu/micro/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct{ *gorm.DB }

func New(user, password, dName string) (storage.Service, error) {
	log.Printf(user)
	log.Printf(password)
	log.Printf(dName)
	b, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return &Mysql{b}, nil
}

func (db *Mysql) Close() error { return db.Close() }

func (db *Mysql) LoadTables() {

	db.AutoMigrate(&model.User{})
	db.AutoMigrate()
}

func (db *Mysql) SaveUser(user model.User) (*model.User, error) {

	// defer db.Close()
	err := db.Create(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (db *Mysql) GetUsers() ([]model.User, error) {
	var users []model.User
	db.Where("is_deleted =?", false).Find(&users)
	return users, nil
}
