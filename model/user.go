package model

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	LastSeen     *time.Time
	AboutMe      string  `gorm:"type:varchar(140)"`
	Avatar       string  `gorm:"type:varchar(200)"`
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

// SetPassword saves user's password by saving its hashed checksum to ensure data security.
func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

// CheckPassword verifies user password.
func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

// GetUserByUsername returns &user if exists.
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserByUsername is used to update user information like last seen, self-introduction etc.
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	return db.Model(user).Update(contents).Error
}

func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

func UpdateAboutMe(username, introduction string) error {
	contents := map[string]interface{}{"about_me": introduction}
	return UpdateUserByUsername(username, contents)
}

func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

//	Follow and Unfollow user, Followers and Following
func (u *User) Follow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	return db.Model(other).Association("Followers").Append(u).Error
}

func (u *User) Unfollow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	return db.Model(other).Association("Followers").Delete(u).Error
}

func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}

func (u *User) FollowingIDs() []int {
	var ids []int
	rows, err := db.Table("follower").Where("follower_id = ?", u.ID).Select("user_id, follower_id").Rows()
	if err != nil {
		log.Println("Counting following error: ", err)
		return ids
	}
	defer rows.Close()

	for rows.Next() {
		var id, followerID int
		_ = rows.Scan(&id, &followerID)
		ids = append(ids, id)
	}
	return ids
}

func (u *User) FollowingCount() int {
	return len(u.FollowingIDs())
}

func (u *User) FollowingPosts() (*[]Post, error) {
	var posts []Post
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", u.FollowingIDs()).Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func (u *User) IsFollowedByUser(username string) bool {
	user, err := GetUserByUsername(username)
	if err != nil {
		log.Printf("User %s does not exist!\n", username)
	}
	ids := user.FollowingIDs()
	for _, id := range ids {
		if id == u.ID {
			return true
		}
	}
	return false
}

func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}

func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	user.SetAvatar(email)
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	// follow self when add user
	return user.FollowSelf()
}

func (u *User) FollowingPostsByPageAndLimit(current, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (current - 1) * limit
	ids := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id in (?)", ids).Count(&total)
	return &posts, total, nil
}
