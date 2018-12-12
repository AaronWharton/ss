package view

import "ss/model"

type IndexView struct {
	BaseView
	model.User
	Posts []model.Post
}

type IVM struct{}

func (IVM) GetView() IndexView {
	u1, u2 := model.User{Username: "Aaron"}, model.User{Username: "Allen"}
	p1, p2 := model.Post{Body: "Today is a good day!", User: u1}, model.Post{Body: "Sure it is!", User: u2}
	return IndexView{BaseView{Title: "Homepage"}, u1, []model.Post{p1, p2}}
}
