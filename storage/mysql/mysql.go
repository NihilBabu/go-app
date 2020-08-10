package mysql

import (
	"github.com/NihilBabu/micro/model"
	"github.com/NihilBabu/micro/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct{ *gorm.DB }

func New(user, password, dName,url string) (storage.Service, error) {
	b, err := gorm.Open("mysql", user+":"+password+"@tcp("+url+")/"+dName+"?charset=utf8&parseTime=True&loc=Local")
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
	err := db.Create(&user).Error
	return &user, err
}

func (db *Mysql) GetUsers() ([]model.User, error) {
	var users []model.User
	err := db.Where("is_deleted =?", false).Find(&users).Error
	return users, err
}

func (db *Mysql) GetUser(userId string) (model.User, error) {
	var user model.User
	err := db.Where("id =? and is_deleted =?", userId,false).Find (&user).Error
	return user, err
}