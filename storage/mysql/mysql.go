package mysql

import (
	"log"

	//mysql driver for gorm

	"github.com/NihilBabu/go-app/model"
	"github.com/NihilBabu/go-app/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct{ db *gorm.DB }

func (p *Mysql) SaveUser(model.User) (model.User, error) {

	return model.User{
		Email: "hlo",
	}, nil

}

func (p *Mysql) Save(url string) (string, error) {
	return url, nil
}

// New returns a mysql backed storage service
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

//Close function for closing mysql connection
func (p *Mysql) Close() error { return p.db.Close() }

// func GetDatabaseConnection() *gorm.DB {

// 	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		log.Panic(err)
// 	} // defer db.Close()
// 	return db

// }
