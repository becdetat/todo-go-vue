package main

import(
	"database/sql"
)

func DeleteTodo( db *sql.DB, id int ) error {
	_, err := db.Exec(
		"DELETE FROM todos WHERE id = $1",
		id )

	return err
}
