package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetTodos( db *sql.DB, c *gin.Context ) {
	todos, err := GetTodos( db )
	if err != nil {
		c.String( http.StatusInternalServerError,
			fmt.Sprintf( "Error getting todos: %q", err ) )
		return
	}
	c.JSON( http.StatusOK, todos )
}
