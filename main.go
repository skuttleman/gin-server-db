package main

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	_ "github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/joho/godotenv/autoload"
	"github.com/skuttleman/gin-server/routes"
)

func main() {
	app := gin.Default()

	routes.Authors(app.Group("/authors"))
	routes.Books(app.Group("/books"))

	app.Run()
}
