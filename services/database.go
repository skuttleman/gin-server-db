package services

import (
  "github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
  "os"
  "database/sql"
  _ "github.com/joho/godotenv/autoload"
	_ "github.com/go-sql-driver/mysql"
  "fmt"
)

var db *sql.DB

func init() {
  db, _ = sql.Open("mysql", os.Getenv("DATABASE_URL"))
}

func Query(fn func(rows *sql.Rows) []gin.H, query string) ([]gin.H, error) {
  rows, err := db.Query(query)
  if err != nil {
    fmt.Println(err.Error())
    return nil, err
  }
  defer rows.Close()
  return fn(rows), nil
}
