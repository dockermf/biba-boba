package dbapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Returns (true, nil) if account with username `username` was found in database;
// otherwise returns (false, nil). If unexpected error happens, (false, error) is returned.
func accountExistsByName(ctx context.Context, db *pgxpool.Pool, username string) (bool, error) {
	var existingAccount string
	err := db.QueryRow(
		ctx,
		"SELECT username FROM Account WHERE username = $1",
		username,
	).Scan(&existingAccount)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		} else {
			return false, fmt.Errorf("Unexpected error: %w", err)
		}
	}
	return true, nil
}

// Returns (true, nil) if account with ID `ID` was found in database; otherwise
// returns (false, nil). If unexpected error happens, (false, error) is returned.
func accountExistsByID(ctx context.Context, db *pgxpool.Pool, ID uint64) (bool, error) {
	var existingAccount string
	err := db.QueryRow(
		ctx,
		"SELECT username FROM Account WHERE account_id = $1",
		ID,
	).Scan(&existingAccount)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		} else {
			return false, fmt.Errorf("Unexpected error: %w", err)
		}
	}
	return true, nil
}
