package view

import (
	"log"
	"ss/model"
)

type LoginView struct {
	BaseView
	Errors []string			// errors information
}

type LVM struct{}

func (LVM) GetView() LoginView {
	v := LoginView{}
	v.SetTitle("Login")
	return v
}

// add login errors information
func (v *LoginView) AddError(err ...string) {
	v.Errors = append(v.Errors, err...)
}

func CheckLogin(username, password string) bool {
	user ,err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Error: ", err)
		return false
	}
	return user.CheckPassword(password)
}
