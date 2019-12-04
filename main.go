package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

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

	router.Use( cors.New( cors.Config {
		AllowOrigins: []string {
			"http://thawing-bayou-17829.herokuapp.com",
			"http://localhost:8080",
		},
	} ) )

	router.Use( static.Serve( "/", static.LocalFile( "dist", true ) ) )

	router.GET( "/todos", func( c *gin.Context ) {
		HandleGetTodos( db, c )
	} )
	router.POST( "/todos", func( c *gin.Context ) {
		HandlePostTodo( db, c )
	} )
	router.PUT( "/todos/:id", func( c *gin.Context ) {
		HandlePutTodo( db, c )
	} )

	router.Run(":" + port)
}

