package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goapp"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	DB = db
	log.Println("Connected to MySQL")
}
