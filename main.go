package main

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
  "github.com/skuttleman/gin-server/routes"
  "database/sql"
  _ "github.com/joho/godotenv/autoload"
  _ "github.com/go-sql-driver/mysql"
  "os"
  "fmt"
)

func main() {
  dburl := os.Getenv("DATABASE_URL")
  fmt.Println(dburl)
	app := gin.Default()
  db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
  fmt.Println(os.Getenv("DATABASE_URL"))
  if err != nil {
    fmt.Println(err)
    return
  }

	app.Use(testMiddleware())

	routes.Authors(app.Group("/authors"), db)
	routes.Books(app.Group("/books"), db)

	app.Run()
}

func testMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("foo", "bar")
		c.Next()
	}
}
