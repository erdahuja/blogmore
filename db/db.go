package db

import (
	"blogmore/utils"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // register postgres driver
	"github.com/joho/godotenv"
)

// EnvVars exposes map of env variables
var EnvVars map[string]string

func init() {
	var err error
	EnvVars, err = godotenv.Read("development.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Env variables are intialized...")
}

// Database exposes db instance and related services
type Database struct {
	DB   *gorm.DB
	Hmac utils.HMAC
}

// New create a connection to database
func New() (*Database, error) {
	user := EnvVars["User"]
	pwd := EnvVars["Password"]
	url := EnvVars["URL"]
	port := EnvVars["Port"]
	dbName := EnvVars["Database"]
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pwd, url, port, dbName)
	db, err := gorm.Open("postgres", dbConnString)
	if err != nil {
		log.Println(dbConnString)
		log.Fatal("Error connecting to database.", err)
		return nil, err
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	hmac := utils.NewHMAC(EnvVars["HMACKey"])
	dBService := Database{
		DB:   db,
		Hmac: hmac,
	}
	return &dBService, nil
}

// First will return first record matched
// if record found, return record
// if any other error, return error with more information
func First(db *gorm.DB, dst interface{}) error {
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
