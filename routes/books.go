package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/skuttleman/gin-server/services"
	"database/sql"
	"fmt"
)

func Books(i *gin.RouterGroup) {
	//C
	i.POST("/", func(c *gin.Context) {
		title := c.PostForm("title")
		if title == "" {
			c.JSON(200, gin.H{
				"success": false,
			})
		} else {
			query := fmt.Sprintf("INSERT INTO books (title) VALUES ('%s')", title)
			_, _ = services.Query(bookProcess, query)
			c.JSON(200, gin.H{
				"success": true,
			})
		}
	})

	//R
	i.GET("/", func(c *gin.Context) {
		rows, _ := services.Query(bookProcess, "SELECT * FROM books")
  	c.JSON(200, gin.H{
			"books": rows,
		})
	})

	//U
	i.PUT("/:id", func(c *gin.Context) {
		title := c.PostForm("title")
		if title =="" {
			c.JSON(200, gin.H{
				"success": false,
			})
		} else {
			query := fmt.Sprintf("UPDATE books SET title='%s' WHERE id=%s", title, c.Param("id"))
			_, _ = services.Query(bookProcess, query)
			c.JSON(200, gin.H{
				"success": true,
			})
		}
	})

	//D
	i.DELETE("/:id", func(c *gin.Context) {
		_, _ = services.Query(bookProcess, "DELETE FROM books WHERE id=" + c.Param("id"))
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	//L
  i.GET("/:id", func(c *gin.Context) {
		rows, _ := services.Query(bookProcess, "SELECT * FROM books WHERE id=" + c.Param("id"))
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
