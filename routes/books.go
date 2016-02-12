package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/skuttleman/gin-server/services"
	"database/sql"
)

func Books(i *gin.RouterGroup) {
	i.GET("/", func(c *gin.Context) {
		rows, _ := services.Query("SELECT * FROM books", bookProcess)
  	c.JSON(200, gin.H{
			"books": rows,
		})
	})
  i.GET("/:id", func(c *gin.Context) {
		rows, _ := services.Query("SELECT * FROM books where id=" + c.Param("id"), bookProcess)
  	c.JSON(200, gin.H{
			"books": rows,
		})
	})
}

func bookProcess(records *sql.Rows) []gin.H {
	items := []gin.H{}
	for records.Next() {
		var id int
		var title string
		records.Scan(&id, &title)
		items = append(items, gin.H{"id": id, "title": title, })
	}
	return items
}
