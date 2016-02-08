package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
  "database/sql"
  // "encoding/json"
  "fmt"
)

func Books(i *gin.RouterGroup, db *sql.DB) {
	i.GET("/", func(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
      fmt.Println(err)
      c.JSON(404, gin.H{
        "error": err.Error(),
      })
      return
    }
    defer rows.Close()
    books := []gin.H{}
    for rows.Next() {
      var id int
      var title string
      err = rows.Scan(&id, &title)
      books = append(books, gin.H{"id": id, "title": title, })
    }
    fmt.Println(books)
  	c.JSON(200, gin.H{
			"books": books,
		})
	})
	i.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"book": c.Param("id"),
		})
	})
}
