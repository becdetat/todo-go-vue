package main

import (
	"database/sql"
)

func GetTodos( db *sql.DB ) ( []*Todo, error ) {
	rows, err := db.Query( "SELECT id, title, position, complete FROM todos" )
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*Todo {}

	for rows.Next() {
		todo := &Todo {}

		if err := rows.Scan( &todo.Id, &todo.Title, &todo.Position, &todo.Complete ); err != nil {
			return nil, err
		}

		todos = append( todos, todo )
	}

	return todos, nil
}
