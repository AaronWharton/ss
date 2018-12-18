package view

import "ss/model"

type IndexView struct {
	BaseView
	Posts []model.Post
	Flash string

	BasePageView
}

type IVM struct{}

func (IVM) GetView(username, flash string, current, limit int) IndexView {
	u, _ := model.GetUserByUsername(username)
	posts, total, _ := u.FollowingPostsByPageAndLimit(current, limit)
	v := IndexView{}
	v.SetTitle("HomePage")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageView(total, current, limit)
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
