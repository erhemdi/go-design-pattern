package objectpool

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDBConnection() *sql.DB {
	return db
}

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/sharing-session")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
}
