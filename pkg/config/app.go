package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// db, err := sql.Open("mysql", "testUser:testUser@tcp(127.0.0.1:3306)/testdb")

	// testUser:testUser@tcp(127.0.0.1:3306)/testdb
	// d, err := sql.Open("mysql", "testUser:testUser@tcp(127.0.0.1:3306)/testdb")
	d, err := gorm.Open("mysql", "francis:francis@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
