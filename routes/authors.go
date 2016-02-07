package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	// "fmt"
)

func Authors(i *gin.RouterGroup) {
	i.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"authors": "list",
		})
	})
	i.GET("/:id", func(c *gin.Context) {
		bar, _ := c.Get("foo")
		c.JSON(200, gin.H{
			"author": c.Param("id"),
			"foo":    bar,
		})
	})
}
