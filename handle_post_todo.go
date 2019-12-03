package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePostTodo( db *sql.DB, c *gin.Context ) {
	var todo Todo
	c.BindJSON( &todo )
	err := AddTodo( db, &todo )
	if err != nil {
		c.String( http.StatusInternalServerError,
			fmt.Sprintf( "Error adding todo: %q", err ) )
		return
	}

	c.String( http.StatusOK, "Success" )
}
