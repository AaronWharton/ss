package main

import (
	"fmt"
	"ss/model"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

func main() {
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username:     "aaron",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "aaron@123.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("aaron@123.com")),
			Posts: []model.Post{
				{Body: "Today is a good day!"},
			},
		},
		{
			Username:     "allen",
			PasswordHash: model.GeneratePasswordHash("123abc"),
			Email:        "allen@456.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("allen@456.com")),
			Posts: []model.Post{
				{Body: "Yes it is!"},
				{Body: "Sun shine is beautiful!"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
