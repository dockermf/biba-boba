package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/dockermf/biba-boba/include/dbapi"
	_ "github.com/gin-gonic/gin"
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
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
