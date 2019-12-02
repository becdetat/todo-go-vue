package main

import (
	"database/sql"
	"log"
)

func MigrateDatabase( db *sql.DB ) {
	if _, err := db.Exec( "CREATE TABLE IF NOT EXISTS ticks (tick timestamp)" ); err != nil {
		log.Fatalf( "Error executing statement: %q", err )
	}

	if _, err := db.Exec( "CREATE TABLE IF NOT EXISTS todos (id int, title text, position int, complete bool)" ); err != nil {
		log.Fatalf( "Error executing statement: %q", err )
	}
}
