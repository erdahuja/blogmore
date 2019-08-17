package services

import (
	"blogmore/db"
	"blogmore/models"
	"blogmore/utils"
	"errors"
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
	err := dbService.First(db, &user)
	return &user, err
}

// ByEmail will lookup user in db for provided email id
// if user found, return user
// if any other error, return error with more information
func (us *UserService) ByEmail(email string) (*models.User, error) {
	var user models.User
	dbService := db.DBService
	db := dbService.DB.Where("email=?", email)
	err := dbService.First(db, &user)
	return &user, err
}

// Login authenticate user with email and password
// if success, return user
// if error, return error with more information
func (us *UserService) Login(email, pwd string) (*models.User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := utils.CompareHashAndPassword([]byte(foundUser.Token), []byte(pwd+db.EnvVars["PwdPepper"])); err != nil {
		return nil, err
	}
	return foundUser, nil
}

// Create adds a new user to db
func (us *UserService) Create(user *models.User) (*models.User, error) {
	hashedBytes, err := utils.GenerateHash([]byte(user.Password + db.EnvVars["PwdPepper"]))
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
// TODO search by email and then delete
// func (us *UserService) Delete(id uint) error {
// 	if id == 0 {
// 		return ErrInvalidID
// 	}
// 	dbService := db.DBService
// 	user := models.User{Model: gorm.Model{ID: id}}
// 	return dbService.DB.Delete(&user).Error
// }
