package mysql

import (
	"fmt"
	"bookshelf-web-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"bookshelf-web-api/domain/model"
)
type Result struct {
	model.BaseModel
	Name string
}
var cate model.Category
var result Result

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
	return db
}
