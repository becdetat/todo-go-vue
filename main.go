package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.Static( "/", "dist" )

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

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

