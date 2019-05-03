package mysql

import (
	"bookshelf-web-api/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func GetDBConn() *gorm.DB {
	dbconf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.DB)
	db, err := gorm.Open("mysql", dbconf)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	// defer db.Close()
	return db
}
