package database

import (
	"database/sql"
	"log"
)

type (
	// Todo define properties of Todoer
	Todo struct {
		ID          int    `json:"id"`
		Creator     int    `json:"creator"`
		Description string `json:"description"`
		CreatedAt   string `json:"createdAt"`
		ModifiedAt  string `json:"modifiedAt"`
	}

	// TodoModel define concrete type of MODEL interface for todo
	TodoModel struct {
		records int
	}
)

// ReadAll read all records from todo table
func (model *TodoModel) ReadAll(id int) []Todo {
	var todos []Todo

	connect(func(db *sql.DB) {

		rows, err := db.Query("select * from todo where creator = ?", id)

		if err != nil {
			log.Panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var todo = Todo{}
			err := rows.Scan(&todo.ID, &todo.Description, &todo.Creator, &todo.CreatedAt, &todo.ModifiedAt)
			if err != nil {
				log.Panic(err)
			}
			todos = append(todos, todo)
		}

		err = rows.Err()

		if err != nil {
			log.Panic(err)
		}
	})

	return todos
}

// Read a record from todo table
func (model *TodoModel) Read(id int) Todo {
	var todo = Todo{}

	connect(func(db *sql.DB) {
		rows, err := db.Query("select * from todo  where id = ?", id)

		if err != nil {
			log.Panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&todo.ID, &todo.Description, &todo.Creator, &todo.CreatedAt, &todo.ModifiedAt)
			if err != nil {
				log.Panic(err)
			}
		}

		err = rows.Err()

		if err != nil {
			log.Panic(err)
		}
	})

	return todo
}

// Add a record to the todo table
func (model *TodoModel) Add(todo Todo) Todo {
	var insertedID int64

	connect(func(db *sql.DB) {
		stmt, err := db.Prepare("INSERT INTO todo(description, creator) VALUES(?, ?)")

		res, err := stmt.Exec(todo.Description, todo.Creator)

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

// Update a record in the todo table
func (model *TodoModel) Update(todo Todo) Todo {
	connect(func(db *sql.DB) {
		stmt, err := db.Prepare("UPDATE todo SET description=? WHERE id=?")

		_, err = stmt.Exec(todo.Description, todo.ID)

		if err != nil {
			log.Panic(err)
		}
	})

	return model.Read(int(todo.ID))
}

// Delete a record from the todo table
func (model *TodoModel) Delete(id int) int {
	var affectedRows int64
	connect(func(db *sql.DB) {
		result, err := db.Exec("delete from todo where id = ?", id)

		if err != nil {
			log.Panic(err)
		}

		affectedRows, _ = result.RowsAffected()
	})

	return int(affectedRows)
}
