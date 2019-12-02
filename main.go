package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
  "github.com/russross/blackfriday"
)

func dbFunc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ExecuteStatement( db, "INSERT INTO ticks VALUES (now())" )
		rows := ExecuteQuery( db, "SELECT tick FROM ticks" )

		defer rows.Close()
		for rows.Next() {
			var tick time.Time
			if err := rows.Scan(&tick); err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error scanning ticks: %q", err))
				return
			}
			c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
			}
		}
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

	router.GET( "/mark", func( c *gin.Context ) {
		c.String( http.StatusOK, string( blackfriday.Run( []byte( "**hi!**" ) ) ) )
	} )

	router.GET("/db", dbFunc(db))

	router.GET( "/todos", func( c *gin.Context ) {
		todos, err := GetTodos( db )
		if err != nil {
			c.String( http.StatusInternalServerError,
				fmt.Sprintf( "Error getting todos", err ) )
			return
		}
		c.JSON( http.StatusOK, todos )
	} )

	router.Run(":" + port)
}

