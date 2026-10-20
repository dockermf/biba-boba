package dbapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Returns true if account with username `username` was found in database;
// otherwise returns false.
func accountExistsByName(ctx context.Context, db *pgxpool.Pool, username string) bool {
	var existingAccount string
	err := db.QueryRow(
		ctx,
		"SELECT username FROM Account WHERE username = $1",
		username,
	).Scan(&existingAccount)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("Bullshit error happened:", err);
		}
		return false
	}
	return true
}

// Returns true if account with ID `ID` was found in database; otherwise
// returns false.
func accountExistsByID(ctx context.Context, db *pgxpool.Pool, ID uint64) bool {
	var existingAccount string
	err := db.QueryRow(
		ctx,
		"SELECT username FROM Account WHERE account_id = $1",
		ID,
	).Scan(&existingAccount)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("Bullshit error happened:", err);
		}
		return false
	}
	return true
}
