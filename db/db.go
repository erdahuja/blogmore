package db

import (
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
	fmt.Println("Creating db instance")
	if err = New(); err != nil {
		log.Fatal("Unable to create db instance", err)
	}
	fmt.Println("Db instance ready")
}

// Database exposes db instance and related services
type Database struct {
	DB *gorm.DB
}

// DBService is database instance
var DBService Database

// New create a connection to database
func New() error {
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
		return err
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		return err
	}
	DBService.DB = db
	return nil
}

// Close closes the db connection
func (dbs *Database) Close() error {
	fmt.Println("closing db")
	return dbs.DB.Close()
}


// First will return first record matched
// if record found, return record
// if any other error, return error with more information
func (dbs *Database) First(db *gorm.DB, dst interface{}) error {
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