package main

import (
	"database/sql"
)

func UpdateTodo( db *sql.DB, id int, todo *Todo ) error {
	_, err := db.Exec(
		"UPDATE todos SET title = $2, position = $3, complete = $4 WHERE id = $1",
		id,
		todo.Title,
		todo.Position,
		todo.Complete )

	return err
}
