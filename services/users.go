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
	// ErrInvalidPassword when password is incorrect
	ErrInvalidPassword = errors.New("invalid password for user")
)

// UserDB is sued to interact with database for user model
type UserDB interface {
	// Methods for querying single users
	ByID(id uint) (*models.User, error)
	ByEmail(email string) (*models.User, error)
	ByRemember(token string) (*models.User, error)

	// Methods for altering users
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error

	// Close db connection
	Close() error

	// Migration helpers
	AutoMigrate()
	DestructiveReset()
}

// UserService are methods that can be operated on user model
// it does work of db interaction
// this is a lower level API which will be used in  controllers for binding rest api and db
type UserService interface {
	UserDB
	Login(email, pwd string) (*models.User, error)
}

type userService struct {
	UserDB
	db *db.Database
}

// Login authenticate user with email and password
// if success, return user
// if error, return error with more information
func (us *userService) Login(email, pwd string) (*models.User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := utils.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(pwd+db.EnvVars["PwdPepper"])); err != nil {
		return nil, ErrInvalidPassword
	}
	return foundUser, nil
}

var _ UserService = &userService{}

// NewUserService creates new user service instance
func NewUserService() (UserService, error) {
	dbService, err := db.New()
	if err != nil {
		return nil, err
	}
	return &userService{
		db: dbService,
	}, nil
}

// AutoMigrate the schema of database if needed
func (us *userService) AutoMigrate() {
	if db.EnvVars["IS_PRODUCTION"] == "true" {
	us.db.DB.AutoMigrate(&models.User{})
	us.db.DB.AutoMigrate(&models.Follow{})
	}
}

// DestructiveReset drops all tables
func (us *userService) DestructiveReset() {
	if db.EnvVars["IS_PRODUCTION"] == "true" {
		us.db.DB.DropTableIfExists(&models.User{})
		us.db.DB.DropTableIfExists(&models.Follow{})
	}
}

// Close closes the db connection
func (us *userService) Close() error {
	return us.db.DB.Close()
}

// ByID will lookup user in db for provided id
// if user found, return user
// if any other error, return error with more information
func (us *userService) ByID(id uint) (*models.User, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}
	var user models.User
	dbi := us.db.DB.Where("id=?", id)
	err := db.First(dbi, &user)
	return &user, err
}

// ByEmail will lookup user in db for provided email id
// if user found, return user
// if any other error, return error with more information
func (us *userService) ByEmail(email string) (*models.User, error) {
	var user models.User
	dbi := us.db.DB.Where("email=?", email)
	err := db.First(dbi, &user)
	return &user, err
}

// ByRemember looks up a user by remember token,
// it will handle hashing for us
func (us *userService) ByRemember(token string) (*models.User, error) {
	hashedToken := us.db.Hmac.Hash(token)
	var user models.User
	dbi := us.db.DB.Where("remember_token_hash=?", hashedToken)
	if err := db.First(dbi, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Create adds a new user to db
func (us *userService) Create(user *models.User) (*models.User, error) {
	hashedBytes, err := utils.GenerateHash([]byte(user.Password + db.EnvVars["PwdPepper"]))
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hashedBytes)
	if user.RememberToken != "" {
		user.RememberTokenHash = us.db.Hmac.Hash(user.RememberToken)
	}
	err = us.db.DB.Create(user).Error
	switch err {
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// Update all fields in a db by PK
func (us *userService) Update(user *models.User) (*models.User, error) {
	if user.RememberToken != "" {
		user.RememberTokenHash = us.db.Hmac.Hash(user.RememberToken)
	}
	err := us.db.DB.Save(user).Error
	switch err {
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// Delete user record by ID
func (us *userService) Delete(id uint) error {
	var user models.User
	if id <= 0 {
		return ErrInvalidID
	}
	user.ID = id
	return us.db.DB.Delete(&user).Error
}
