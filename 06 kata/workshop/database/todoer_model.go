package database

import (
	"database/sql"
	"log"
)

type (
	// Todoer define properties of Todoer
	Todoer struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		CreatedAt  string `json:"createdAt"`
		ModifiedAt string `json:"modifiedAt"`
	}

	// TodoerModel define concrete type of MODEL interface for todoer
	TodoerModel struct{}
)

// Read a record from todoer table
func (model *TodoerModel) Read(id int) Todoer {
	var todoer = Todoer{}

	connect(func(db *sql.DB) {
		rows, err := db.Query("select * from todoer  where id = ?", id)

		if err != nil {
			log.Panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&todoer.ID, &todoer.Username, &todoer.CreatedAt, &todoer.ModifiedAt)
			if err != nil {
				log.Panic(err)
			}
		}

		err = rows.Err()

		if err != nil {
			log.Panic(err)
		}
	})

	return todoer
}

// Add a record to the todoer table
func (model *TodoerModel) Add(username string) Todoer {
	var insertedID int64

	connect(func(db *sql.DB) {
		stmt, err := db.Prepare("INSERT INTO todoer(username) VALUES(?)")

		res, err := stmt.Exec(username)

		if err != nil {
			log.Panic(err)
		}

		lastID, err := res.LastInsertId()

		if err != nil {
			log.Panic(err)
		}

		insertedID = lastID
	})

	return model.Read(int(insertedID))
}

// Update a record to the todoer table
func (model *TodoerModel) Update(id int, username string) Todoer {
	connect(func(db *sql.DB) {
		stmt, err := db.Prepare("UPDATE todoer SET username=? WHERE id=?")

		_, err = stmt.Exec(username, id)

		if err != nil {
			log.Panic(err)
		}
	})

	return model.Read(int(id))
}

// Delete a record from the todoer table
func (model *TodoerModel) Delete(id int) int {
	var affectedRows int64
	connect(func(db *sql.DB) {
		result, err := db.Exec("delete from todoer where id = ?", id)

		if err != nil {
			log.Panic(err)
		}

		affectedRows, _ = result.RowsAffected()
	})

	return int(affectedRows)
}
