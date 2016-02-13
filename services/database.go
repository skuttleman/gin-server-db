package services

import (
	"database/sql"
	"fmt"
	"github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/gin-gonic/gin"
	_ "github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
	_ "github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/joho/godotenv/autoload"
	"os"
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
