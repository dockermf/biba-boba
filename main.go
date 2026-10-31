package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dockermf/biba-boba/include/dbapi"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// postgres://username:password@host:5432/database_name
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err)
	}
	// dirty example setup
	router := gin.Default()
	router.GET("/api/create-account", func(c *gin.Context) {
		query := c.Request.URL.Query()
		username := query.Get("username")
		password := query.Get("password")
		var message string
		var retcode int
		// hardcoded regular type lol
		if err := dbapi.CreateAccount(context.Background(), pool, dbapi.AccountType_Regular, username, password); err != nil {
			retcode = 200
			message = err.Error()
		} else {
			retcode = 201
			message = "OK"
		}
		c.JSON(retcode, gin.H{
			"response": message,
		})
	})
	router.Run("0.0.0.0:8080")
	defer pool.Close()
}
