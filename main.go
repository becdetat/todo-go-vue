package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func handleGetTodos( db *sql.DB, c *gin.Context ) {
	todos, err := GetTodos( db )
	if err != nil {
		c.String( http.StatusInternalServerError,
			fmt.Sprintf( "Error getting todos: %q", err ) )
		return
	}
	c.JSON( http.StatusOK, todos )
}

func handlePostTodos( db *sql.DB, c *gin.Context ) {
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

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

  db, err := sql.Open( "postgres", os.Getenv( "DATABASE_URL" ) )
  if err != nil {
    log.Fatalf( "Error opening database: %q", err )
  }

  MigrateDatabase( db )

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET( "/todos", func( c *gin.Context ) {
		handleGetTodos( db, c )
	} )
	router.POST( "/todos", func( c *gin.Context ) {
		handlePostTodos( db, c )
	} )

	router.Run(":" + port)
}

