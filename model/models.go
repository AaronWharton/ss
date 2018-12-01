package model

type User struct {
	UserName string
}

type Post struct {
	Body string
	User
}
