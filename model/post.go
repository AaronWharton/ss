package model

import (
	"time"
)

type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User // TODO: why can not use `User` only?
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

func GetPostByUserIDPageAndLimit(id, current, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (current - 1) * limit
	if err := db.Preload("User").Order("timestamp desc").Where("user_id=?", id).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id=?", id).Count(&total)
	return &posts, total, nil
}

func GetPostByPageAndLimit(current, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (current - 1) * limit
	if err := db.Preload("User").Offset(offset).Limit(limit).Order("timestamp desc").Find(&posts).Error; err != nil {
		return nil, total, nil
	}
	db.Model(&Post{}).Count(&total)
	return &posts, total, nil
}

func (p *Post) FormattedTimeAgo() string {
	return FromTime(*p.Timestamp)
}
