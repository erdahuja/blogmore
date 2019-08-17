package models

import (
	"blogmore/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	AutoMigrate()
}

// User model represented in db
type User struct {
	gorm.Model
	Username   string `gorm:"unique_index;not null"`
	Email      string `gorm:"unique_index;not null"`
	Token      string `gorm:"not null"`
	Bio        string
	Image      string
	Followers  []Follow `gorm:"foreignkey:FollowingID"`
	Followings []Follow `gorm:"foreignkey:FollowedByID"`
}

// Follow - jon is folowwing susan
// FollowedByID-jon
// FollowingID-susan
type Follow struct {
	gorm.Model
	FollowingID  uint `gorm:"primary_key" sql:"type:int not null"`
	FollowedByID uint `gorm:"primary_key" sql:"type:int not null"`
}

// AutoMigrate the schema of database if needed
func AutoMigrate() {
	dbService := db.New()
	dbService.Db.DropTableIfExists(&User{})
	dbService.Db.DropTableIfExists(&Follow{})
	dbService.Db.AutoMigrate(&User{})
	dbService.Db.AutoMigrate(&Follow{})
}
