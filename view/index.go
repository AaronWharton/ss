package view

import "ss/model"

type IndexView struct {
	BaseView
	model.User
	Posts []model.Post
}

type IVM struct{}

func (IVM) GetView() IndexView {
	u1, _ := model.GetUserByUsername("aaron")
	posts, _ := model.GetPostsByUserID(u1.ID)
	return IndexView{BaseView{Title: "Homepage"}, *u1, *posts}
}
