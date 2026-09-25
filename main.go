package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Account struct {
	ID       uint64
	Login    string
	Password string
	Balance  float64
}

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/testdb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed connecting to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var login string
	var balance float64
	sql_query := "select login, balance from users"

	// Gets rows and stores them in rows var
	rows, _ := conn.Query(context.Background(), sql_query)
	// For each row in rows var, scan results into login and balance, then run func()
	_, err = pgx.ForEachRow(rows, []any{&login, &balance}, func() error {
		fmt.Fprintf(os.Stdout, "%s %f\n", login, balance)
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "ForEachRow failed: %v\n", err)
		os.Exit(1)
	}
}
