package view

import (
	"ss/model"
)

type ExploreView struct {
	BaseView
	BasePageView
	Posts []model.Post
}

type EVM struct{}

func (EVM) GetView(username string, current, limit int) ExploreView {
	posts, total, _ := model.GetPostByPageAndLimit(current, limit)
	v := ExploreView{}
	v.SetTitle("Explore")
	v.SetCurrentUser(username)
	v.SetBasePageView(total, current, limit)
	v.Posts = *posts
	return v
}
