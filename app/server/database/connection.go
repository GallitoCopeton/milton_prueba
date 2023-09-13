package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Connection struct {
	DriverName         string
	DataSourceName     string
	MaxIdleConnections int
	MaxOpenConnections int
	MaxConnIdleTime    time.Duration
	MaxConnMaxLifetime time.Duration
	Logger             *log.Logger
}

type Database struct {
	*Connection
	*sqlx.DB
}

func Create(createDB *Connection) *Database {
	db, err := sqlx.Open(createDB.DriverName, createDB.DataSourceName)
	if err != nil {
		fmt.Println(createDB)
		fmt.Println(err)
		db.Close()
		createDB.Logger.Fatalf("Could not connect to database %v", err)
	}

	db.SetMaxIdleConns(createDB.MaxIdleConnections)
	db.SetMaxOpenConns(createDB.MaxOpenConnections)
	db.SetConnMaxIdleTime(createDB.MaxConnIdleTime)
	db.SetConnMaxLifetime(createDB.MaxConnMaxLifetime)


	createDB.Logger.Println("Database connection is now open.git.")

	return NewDB(db, createDB)
}

func NewDB(DB *sqlx.DB, conn *Connection) *Database {
	return &Database{DB: DB, Connection: conn}
}
