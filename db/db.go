package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
}

// Database exposes db instance and related services
type Database struct {
	Db *gorm.DB
}

// New create a connection to database
func New() *Database {
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
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	return &Database{
		Db: db,
	}
}

// Close closes the db connection
func (dbs *Database) Close() error {
	return dbs.Db.Close()
}
