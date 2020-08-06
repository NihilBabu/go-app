package storage

import (
	"log"

	"github.com/NihilBabu/go-app/model"
	"github.com/jinzhu/gorm"

	//mysql driver for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Service interface is ther service connection for all storage
type Service interface {
	SaveUser(model.User) (model.User, error)
	Save(string) (string, error)

	Close() error
}

//GetDatabaseConnection is a old implementation depricated
func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	} // defer db.Close()
	return db

}
