package main

import (
	"database/sql"
)

func MigrateDatabase( db *sql.DB ) {
	ExecuteStatement( db, "CREATE TABLE IF NOT EXISTS ticks (tick timestamp)" )
	ExecuteStatement( db, "CREATE TABLE IF NOT EXISTS todos (id int, title text, position int, complete bool)" )
}
