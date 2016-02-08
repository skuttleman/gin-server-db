package routes

import (
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
  "database/sql"
  // "encoding/json"
  "fmt"
)

func Authors(i *gin.RouterGroup, db *sql.DB) {
	i.GET("/", func(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM authors")
    if err != nil {
      fmt.Println(err)
      c.JSON(404, gin.H{
        "error": err.Error(),
      })
      return
    }
    defer rows.Close()
    authors := []gin.H{}
    for rows.Next() {
      var id int
      var first_name string
      var last_name string
      err = rows.Scan(&id, &first_name, &last_name)
      authors = append(authors, gin.H{"id": id, "last_name": last_name, "first_name": first_name, })
    }
    fmt.Println(authors)
  	c.JSON(200, gin.H{
			"authors": authors,
		})
	})
  i.GET("/:id", func(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM authors where id=?", c.Param("id"))
    if err != nil {
      fmt.Println(err)
      c.JSON(404, gin.H{
        "error": err.Error(),
      })
      return
    }
    defer rows.Close()
    authors := []gin.H{}
    for rows.Next() {
      var id int
      var first_name string
      var last_name string
      err = rows.Scan(&id, &first_name, &last_name)
      authors = append(authors, gin.H{"id": id, "last_name": last_name, "first_name": first_name, })
    }
    fmt.Println(authors)
  	c.JSON(200, gin.H{
			"authors": authors,
		})
	})
}
