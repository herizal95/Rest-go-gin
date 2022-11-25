package services

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	USER := "herizal"
	PASS := "herizal123"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "pustaka_api"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
