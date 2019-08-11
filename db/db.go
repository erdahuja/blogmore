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
	envVars, err := godotenv.Read("../development.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Database exposes db instance and related services
type Database struct {
	db *gorm.DB
}

// New create a connection to database
func New() *Database {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", envVars["User"], envVars["Password"], envVars["URL"], envVars["Port"], envVars["Database"])
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
		db: db,
	}
}

// Close closes the db connection
func (dbs *Database) Close() error {
	return dbs.db.Close()
}
