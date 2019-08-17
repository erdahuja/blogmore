package services

import (
	"blogmore/db"
	"blogmore/models"
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound when resource is nnot found in db
	ErrNotFound = errors.New("user model: resource not found")
	// ErrInvalidID when resource is nnot found in db
	ErrInvalidID = errors.New("invalid id for user record")
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
	dbService := db.DBService
	db := dbService.DB.Where("id=?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail will lookup user in db for provided email id
// if user found, return user
// if any other error, return error with more information
func (us *UserService) ByEmail(email string) (*models.User, error) {
	var user models.User
	dbService := db.DBService
	db := dbService.DB.Where("email=?", email)
	err := first(db, &user)
	return &user, err
}

// first will return first record matched
// if record found, return record
// if any other error, return error with more information
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	switch err {
	case nil:
		return nil
	case err:
		return err
	default:
		return err
	}
}

// Create adds a new user to db
func (us *UserService) Create(user *models.User) (*models.User, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password+db.EnvVars["PwdPepper"]), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Token = string(hashedBytes)
	dbService := db.DBService
	err = dbService.DB.Create(user).Error
	switch err {
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// Update all fields in a db by PK
func (us *UserService) Update(user *models.User) (*models.User, error) {
	dbService := db.DBService
	err := dbService.DB.Save(user).Error
	switch err {
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// Delete user record by ID
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	dbService := db.DBService
	user := models.User{Model: gorm.Model{ID: id}}
	return dbService.DB.Delete(&user).Error
}
