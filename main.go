package main

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v5"
    _ "github.com/gin-gonic/gin"
    _ "github.com/dockermf/biba-boba/include/dbapi"
)

const (
    DB_USERNAME = "postgres"
    DB_PASSWORD = "1234"
    DB_HOST     = "localhost"
    DB_PORT     = "5432"
    DB_NAME     = "testdb"
)

func main() {
    // urlExample := "postgres://username:password@localhost:5432/database_name"
    url := "postgres://" + DB_USERNAME + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME
    conn, err := pgx.Connect(context.Background(), url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed connecting to database: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close(context.Background())
}
