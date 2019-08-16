package services

import (
	"blogmore/db"
	"blogmore/models"
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	// ErrNotFound when resource is nnot found in db
	ErrNotFound = errors.New("user model: resource not found")
	// InvalidId when resource is nnot found in db
	InvalidId = errors.New("invalid id for user record")
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
	defer dbService.Close()
	db := dbService.Db.Where("id=?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail will lookup user in db for provided email id
// if user found, return user
// if any other error, return error with more information
func (us *UserService) ByEmail(email string) (*models.User, error) {
	var user models.User
	dbService := db.New()
	defer dbService.Close()
	db := dbService.Db.Where("email=?", email)
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
	dbService := db.New()
	err := dbService.Db.Create(user).Error
	defer dbService.Close()
	switch err {
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// Update all fields in a db by PK
func (us *UserService) Update(user *models.User) (*models.User, error) {
	dbService := db.New()
	err := dbService.Db.Save(user).Error
	defer dbService.Close()
	switch err {
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// Delete user record by ID
func (us *UserService) Delete(id uint) (error) {
	if id === 0 {
		return nil, InvalidId
	}
	dbService := db.New()
	defer dbService.Close()
	user := models.User{Model: gorm.Model{ID: id}}
	fmt.Println("Deleting user ecord with id ", id, user)
	return dbService.Db.Delete(&user).Error
}
