package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleDeleteTodo( db *sql.DB, c *gin.Context ) {
	id, err := strconv.Atoi( c.Param( "id" ) )
	if err != nil {
		c.String( http.StatusInternalServerError,
			fmt.Sprintf( "Error converting id to integer: %q", err ) )
		return
	}

	if err = DeleteTodo( db, id ); err != nil {
		c.String( http.StatusInternalServerError,
			fmt.Sprintf( "Error deleting todo: %q", err ) )
		return
	}

	c.String( http.StatusOK, "Success" )
}
