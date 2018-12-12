package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"ss/config"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connect to db...")
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		log.Fatal("++++++", err)
	}

	db.SingularTable(true)
	return db
}


