package main

import (
	"database/sql"
	"log"
)

func ExecuteQuery( db *sql.DB, query string ) *sql.Rows {
	if rows, err := db.Query( query ); err != nil {
		log.Panicf( "Error executing query: %q", err )
		return nil
	} else {
		return rows
	}
}
