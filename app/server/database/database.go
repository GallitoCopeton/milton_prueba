package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var db *Database
var once sync.Once

func Connect() *Database {
	// Fetch the database connection parameters from environment variables
	dbUsername := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	// Create connection URL
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	conn := Connection{
		DriverName:         "mysql",
		DataSourceName:     dbSource,
		MaxIdleConnections: 1,
		MaxOpenConnections: 10,
		MaxConnIdleTime:    30 * time.Second,
		MaxConnMaxLifetime: 60 * time.Second,
		Logger:             log.New(os.Stdout, "", log.Ltime),
	}

	// Run this code once only, no matter how many times this function is called
	once.Do(func() {
		db = Create(&conn)
	})

	return db
}
