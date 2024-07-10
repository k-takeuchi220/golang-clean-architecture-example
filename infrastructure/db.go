package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewDB() (*sql.DB, error) {
	dbUser := "user"
	dbPassword := "password"
	dbName := "exampledb"
	dbHost := "127.0.0.1"
	dbPort := "3306"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return db, nil
}
