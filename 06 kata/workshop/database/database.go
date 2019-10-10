package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // import for side effect only
)

type (
	// TODOER implement Todoer data model method signatures
	TODOER interface {
		Add(string) Todoer
		Read(int) Todoer
		Update(int, string) Todoer
		Delete(int) int
	}

	// TODO implement Todo data model method signatures
	TODO interface {
		Add(Todo) Todo
		ReadAll(int) []Todo
		Read(int) Todo
		Update(Todo) Todo
		Delete(int) int
	}
)

// Connect method initialise connection to db
func connect(modelFunc func(db *sql.DB)) {
	log.Print("Connect to database")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todo")

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	log.Print("Successfully connected to MySQL database")

	modelFunc(db)
}
