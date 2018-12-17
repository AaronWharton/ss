package view

import "ss/model"

type ProfileView struct {
	BaseView
	Posts       []model.Post
	ProfileUser model.User
	Editable bool
}

type PVM struct{}

func (PVM) GetView(sUser, pUser string) (ProfileView, error) {
	v := ProfileView{}
	v.SetTitle("Profile")
	user, err :=model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}

	posts, _ := model.GetPostsByUserID(user.ID)
	v.ProfileUser = *user
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	v.Editable = sUser == pUser
	return v, nil
}
