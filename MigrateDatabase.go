package main

import (
	"database/sql"
)

func MigrateDatabase( db *sql.DB ) {
	ExecuteStatement( db, "CREATE TABLE IF NOT EXISTS ticks (tick timestamp)" )
	ExecuteStatement( db, "CREATE TABLE IF NOT EXISTS todos (id int, title string, position int, complete bool)" )
}
