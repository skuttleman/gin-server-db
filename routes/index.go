package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
  "database/sql"
)

func Index(i *gin.RouterGroup, db *sql.DB) {
	i.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"main": "list",
		})
	})
	i.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"main": c.Param("id"),
			// "foo": c.Get("foo"),
		})
	})
}
