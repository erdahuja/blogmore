package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // register postgres driver
	"github.com/joho/godotenv"
)

var envVars map[string]string

func init() {
	var err error
	envVars, err = godotenv.Read("development.env")
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
	user := envVars["User"]
	pwd := envVars["Password"]
	url := envVars["URL"]
	port := envVars["Port"]
	dbName := envVars["Database"]
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
