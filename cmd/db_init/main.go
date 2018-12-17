package main

import (
	"log"
	"ss/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	model.AddUser("aaron", "abc123", "aaron@123.com")
	model.AddUser("allen", "123abc", "allen@456.com")

	u1, _ := model.GetUserByUsername("aaron")
	u1.CreatePost("What a nice day!")
	model.UpdateAboutMe(u1.Username, `Love listening to the music!`)

	u2, _ := model.GetUserByUsername("allen")
	u2.CreatePost("Yes it is!")
	u2.CreatePost("How do you do?")

	u1.Follow(u2.Username)
}
