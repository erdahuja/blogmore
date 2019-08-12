package services

import (
	"blogmore/db"
	"blogmore/models"
	"errors"
)

var (
	// ErrNotFound when resource is nnot found in db
	ErrNotFound = errors.New("user model: resource not found")
)

// UserService are methods that can be operated on user model
// it does work of db interaction
// this is a lower level API which will be used in  controllers for binding rest api and db
type UserService struct{}

// ByID will lookup user in db for provided id
// if user found, return user
// if any other error, return error with more information
func (us *UserService) ByID(id uint) (*models.User, error) {
	var user models.User
	dbService := db.New()
	err := dbService.db.Where("id=?", id).First(&user).Error
	defer dbService.Close()
	switch err {
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

// Create adds a new user to db
func (us *UserService) Create(user *models.User) (*User, error) {
	dbService := db.New()
	err := dbService.db.Create(user).Error
	defer dbService.Close()
	switch err {
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

// Update all fields in a db by PK
func (us *UserService) Update(user *models.User) (*User, error) {
	dbService := db.New()
	err := dbService.db.Save(user).Error
	defer dbService.Close()
	switch err {
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}
