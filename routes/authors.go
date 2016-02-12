package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/skuttleman/gin-server/services"
	"database/sql"
)

func Authors(i *gin.RouterGroup) {
	i.GET("/", func(c *gin.Context) {
		rows, _ := services.Query("SELECT * FROM authors", authorProcess)
  	c.JSON(200, gin.H{
			"authors": rows,
		})
	})
  i.GET("/:id", func(c *gin.Context) {
		rows, _ := services.Query("SELECT * FROM authors where id=" + c.Param("id"), authorProcess)
  	c.JSON(200, gin.H{
			"authors": rows,
		})
	})
}

func authorProcess(records *sql.Rows) []gin.H {
	items := []gin.H{}
	for records.Next() {
		var id int
		var first_name string
		var last_name string
		records.Scan(&id, &first_name, &last_name)
		items = append(items, gin.H{"id": id, "first_name": first_name, "last_name": last_name, })
	}
	return items
}
