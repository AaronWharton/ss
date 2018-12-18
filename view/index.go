package view

import "ss/model"

type IndexView struct {
	BaseView
	Posts []model.Post
	Flash string
}

type IVM struct{}

func (IVM) GetView(username, flash string) IndexView {
	u, _ := model.GetUserByUsername(username)
	posts, _ := u.FollowingPosts()
	v := IndexView{BaseView{Title: "HomePage"}, *posts, flash}
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
