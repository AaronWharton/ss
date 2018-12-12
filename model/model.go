package model

type User struct {
	Username string
}

type Post struct {
	Body string
	User
}
