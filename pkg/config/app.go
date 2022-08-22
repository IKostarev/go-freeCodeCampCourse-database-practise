package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "codecamptestuser:codecamptestuser@/codecamptestdb?charset=utf8&parseTime=True&loc=Local")  // "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	if err != nil {
		log.Fatalf("Have error - %s is connect db", err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
