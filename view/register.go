package view

import (
	"fmt"
	"ss/model"
)

type RegisterView struct {
	LoginView
}

type RVM struct{}

func (RVM) GetView() RegisterView {
	v := RegisterView{}
	v.SetTitle("Register")
	return v
}

func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		fmt.Println("username doesn't exist!")
		return true
	}

	return false
}

func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}