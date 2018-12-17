package view

import "ss/model"

type ProfileEditView struct {
	LoginView
	ProfileUser model.User
}

type PEV struct {}

func (PEV) GetView(username string) ProfileEditView {
	v := ProfileEditView{}
	u, _ := model.GetUserByUsername(username)
	v.SetTitle("Profile Edit")
	v.SetCurrentUser(username)
	v.ProfileUser = *u
	return v
}

func UpdateAboutMe(username, introduction string) error {
	return model.UpdateAboutMe(username, introduction)
}
