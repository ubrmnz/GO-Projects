package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var ( 
	db *gorm.DB
)

func Connect() {
	conn, err := gorm.Open('mysql', 'srx:srx/bookstore?charset=utf8&parseTime=True&loc=Local')

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
