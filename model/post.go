package model

import "time"

type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User	// TODO: why can not use `User` only?
	Body      string     `gorm:"type:varchar(180)"`
	Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}

func GetPostsByUserID(id int) (*[]Post, error) {
	var posts []Post
	if err := db.Preload("User").Order("timestamp desc").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}
