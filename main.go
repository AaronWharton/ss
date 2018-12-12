package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"ss/controller"
	"ss/model"
)

func main() {
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	controller.Start()
	_ = http.ListenAndServe(":8888", nil)
}
