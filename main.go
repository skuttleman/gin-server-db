package main

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
  "github.com/skuttleman/gin-server/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := gin.Default()

	routes.Authors(app.Group("/authors"))
	routes.Books(app.Group("/books"))

	app.Run()
}
