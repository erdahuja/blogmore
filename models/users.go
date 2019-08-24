package models

import (
	"github.com/jinzhu/gorm"
)

// User model represented in db
type User struct {
	gorm.Model
	Username          string `gorm:"unique_index;not null"`
	Email             string `gorm:"unique_index;not null"`
	Password          string `gorm:"-"`
	PasswordHash      string `gorm:"not null"`
	RememberToken     string `gorm:"-"`
	RememberTokenHash string `gorm:"unique_index;not null"`
	Bio               string
	Image             string
	Followers         []Follow `gorm:"foreignkey:FollowingID"`
	Followings        []Follow `gorm:"foreignkey:FollowedByID"`
}

// Follow - jon is folowwing susan
// FollowedByID-jon
// FollowingID-susan
type Follow struct {
	gorm.Model
	FollowingID  uint `gorm:"primary_key" sql:"type:int not null"`
	FollowedByID uint `gorm:"primary_key" sql:"type:int not null"`
}
