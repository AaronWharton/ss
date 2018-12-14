package view

import "ss/model"

type IndexView struct {
	BaseView
	model.User
	Posts []model.Post
}

type IVM struct{}

func (IVM) GetView(username string) IndexView {
	u1, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(u1.ID)
	v := IndexView{BaseView{Title: "Homepage"}, *u1, *posts}
	v.SetCurrentUser(username)
	return v
}
