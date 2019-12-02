package main

import (
	"database/sql"
	"log"
)

func ExecuteStatement( db *sql.DB, statement string ) {
	if _, err := db.Exec( statement ); err != nil {
		log.Panicf( "Error executing statement: %q", err )
	}
}

