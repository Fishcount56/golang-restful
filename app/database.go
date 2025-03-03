package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/restful-golang-fiber")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
