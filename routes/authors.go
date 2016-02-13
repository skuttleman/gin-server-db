package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/skuttleman/gin-server/services"
	"database/sql"
	"fmt"
)

func Authors(i *gin.RouterGroup) {
	//C
	i.POST("/", func(c *gin.Context) {
		first_name, last_name := c.PostForm("first_name"), c.PostForm("last_name")
		if first_name == "" || last_name == "" {
			c.JSON(200, gin.H{
				"success": false,
			})
		} else {
			query := fmt.Sprintf("INSERT INTO authors (first_name, last_name) VALUES ('%s', '%s')", first_name, last_name)
			_, _ = services.Query(authorProcess, query)
			c.JSON(200, gin.H{
				"success": true,
			})
		}
	})

	//R
	i.GET("/", func(c *gin.Context) {
		rows, _ := services.Query(authorProcess, "SELECT * FROM authors")
  	c.JSON(200, gin.H{
			"authors": rows,
		})
	})

	//U
	i.PUT("/:id", func(c *gin.Context) {
		first_name, last_name := c.PostForm("first_name"), c.PostForm("last_name")
		if first_name == "" && last_name == "" {
			c.JSON(200, gin.H{
				"success": false,
			})
		} else {
			var query string
			if first_name == "" {
				query = fmt.Sprintf("UPDATE authors SET last_name='%s' WHERE id=%s", last_name, c.Param("id"))
			} else if last_name == "" {
				query = fmt.Sprintf("UPDATE authors SET first_name='%s' WHERE id=%s", first_name, c.Param("id"))
			} else {
				query = fmt.Sprintf("UPDATE authors SET first_name='%s', last_name='%s' WHERE id=%s", first_name, last_name, c.Param("id"))
			}
			_, _ = services.Query(authorProcess, query)
			c.JSON(200, gin.H{
				"success": true,
			})
		}
	})

	//D
	i.DELETE("/:id", func(c *gin.Context) {
		_, _ = services.Query(authorProcess, "DELETE FROM authors WHERE id=" + c.Param("id"))
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	//L
  i.GET("/:id", func(c *gin.Context) {
		rows, _ := services.Query(authorProcess, "SELECT * FROM authors WHERE id=" + c.Param("id"))
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
