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

	corsConfig := cors.DefaultConfig()
	router.Use( cors.New( corsConfig ) )

	router.Use( static.Serve( "/", static.LocalFile( "dist", true ) ) )

	router.GET( "/api/v1/todos", func( c *gin.Context ) {
		HandleGetTodos( db, c )
	} )
	router.POST( "/api/v1/todos", func( c *gin.Context ) {
		HandlePostTodo( db, c )
	} )
	router.PUT( "/api/v1/todos/:id", func( c *gin.Context ) {
		HandlePutTodo( db, c )
	} )
	router.DELETE( "/api/v1/todos/:id", func( c *gin.Context ) {
		HandleDeleteTodo( db, c )
	} )

	router.Run(":" + port)
}

