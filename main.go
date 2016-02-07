package main

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/skuttleman/gin-server/routes"
)

func main() {
	app := gin.Default()

	app.Use(testMiddleware())

	routes.Authors(app.Group("/authors"))
	routes.Books(app.Group("/books"))
	app.Run()
}

func testMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("foo", "bar")
		c.Next()
	}
}
