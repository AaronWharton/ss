package view

import "ss/model"

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IVM struct{}

func (IVM) GetVM() IndexViewModel {
	u1, u2 := model.User{UserName: "Aaron"}, model.User{UserName: "Allen"}
	posts := []model.Post{
		{User: u1, Body: "Today is a good day!"},
		{User: u2, Body: "Yes it is!"},
	}

	return IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
}
