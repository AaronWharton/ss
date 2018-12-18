package view

import "ss/model"

type ProfileView struct {
	BaseView
	BasePageView
	Posts          []model.Post
	ProfileUser    model.User
	Editable       bool
	IsFollowed     bool
	FollowersCount int
	FollowingCount int
}

type PVM struct{}

func (PVM) GetView(sUser, pUser string, current, limit int) (ProfileView, error) {
	v := ProfileView{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, total, _ := model.GetPostByUserIDPageAndLimit(u.ID, current, limit)
	v.ProfileUser = *u
	v.Editable = sUser == pUser
	v.SetBasePageView(total, current, limit)

	if !v.Editable {
		v.IsFollowed = u.IsFollowedByUser(sUser)
	}
	v.FollowersCount = u.FollowersCount()
	v.FollowingCount = u.FollowingCount()

	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}

func Follow(sUser, pUser string) error {
	user, err := model.GetUserByUsername(sUser)
	if err != nil {
		return err
	}

	return user.Follow(pUser)
}

func Unfollow(sUser, pUser string) error {
	user, err := model.GetUserByUsername(sUser)
	if err != nil {
		return err
	}

	return user.Unfollow(pUser)
}
