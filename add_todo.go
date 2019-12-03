package main

import (
	"database/sql"
)

func AddTodo( db *sql.DB, todo *Todo ) error {
	_, err := db.Exec(
		"INSERT INTO todos( id, title, position, complete ) VALUES( $1, $2, $3, $4 )",
		todo.Id,
		todo.Title,
		todo.Position,
		todo.Complete )

	return err
}
