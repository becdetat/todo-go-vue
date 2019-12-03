package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlePutTodo( db *sql.DB, c *gin.Context ) {
	id, err := strconv.Atoi( c.Param( "id" ) )
	if err != nil {
		c.String( http.StatusInternalServerError,
			fmt.Sprintf( "Error converting id to integer: %q", err ) )
		return
	}

	var todo Todo
	c.BindJSON( &todo )

	if err := UpdateTodo( db, id, &todo ); err != nil {
		c.String( http.StatusInternalServerError,
			fmt.Sprintf( "Error updating todo: %q", err ) )
		return
	}

	c.String( http.StatusOK, "Success" )
}
