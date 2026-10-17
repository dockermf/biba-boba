// Package dbapi is used to talk with Postgres
package dbapi // import "github.com/dockermf/biba-boba/include/dbapi"

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountType int32

// Enum
const (
	// Regular account enum type
	AccountType_Regular AccountType = 0

	// Admin account enum type
	AccountType_Admin AccountType = 1
)

// Maps for account types
var (
	AccountType_Name = map[AccountType]string{
		0: "regular",
		1: "admin",
	}
	AccountType_Value = map[string]AccountType{
		"regular": 0,
		"admin":   1,
	}
)

// Account is a type used to represent database Account table.
type Account struct {
	// Primary key of the table (PK).
	//
	// Represents `account_id` column.
	ID uint64

	// AccountType enum.
	AccountType AccountType

	// Username is the publicly visible identifier of the account. Used for login.
	//
	// Represents `username` column.
	Username string

	// Password is the hash of account's password.
	//
	// Represents `password` column.
	Password string

	// Balance is account's current balance. Can't be negative.
	//
	// Represents `balance` column.
	Balance float64
}

// ERRORS
var (
	ErrInvalidInput    = errors.New("username or password is empty")
	ErrUsernameTaken   = errors.New("username already taken")
	ErrAccountNotFound = errors.New("account not found")
)

// Insert new row into `Account` table.
func CreateAccount(ctx context.Context, db *pgxpool.Pool, accountType AccountType, username, password string) error {
	if username == "" || password == "" {
		return ErrInvalidInput
	}

	typeName, ok := AccountType_Name[accountType]
	if !ok {
		return fmt.Errorf("invalid account type: %v", accountType)
	}

	var existingAccount string
	err := db.QueryRow(
		ctx,
		"SELECT username FROM Account WHERE username = $1",
		username,
	).Scan(&existingAccount)

	if err == nil {
		return ErrUsernameTaken
	}

	// Проверяем, что запрос завершился именно с ErrNoRows
	if !errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("Expected error pgx.ErrNoRows, got: %w", err)
	}

	// TODO: hash password
	_, err = db.Exec(
		ctx,
		"INSERT INTO Account (account_type, username, password) VALUES ($1, $2, $3)",
		typeName,
		username,
		password,
	)
	if err != nil {
		return fmt.Errorf("Failed to create new account: %w", err)
	}

	return nil
}

// Delete row from `Account` table.
func RemoveAccount(conn *pgx.Conn, a Account) error {
	/* 1. Query for row with id a.ID
	 * 2. If found, remove, else return error
	 */
	return nil
}

// Update `username` column with `new_username`.
func SetUsername(conn *pgx.Conn, a Account, new_username string) error {
	/* 1. Check if row with a.ID exists; if it doesn't, return error
	   2. Check if current login and provided login differ
	   3. If they differ, update; else return error
	*/
	return nil
}

// Update `password` column with `password`.
func SetPassword(conn *pgx.Conn, a Account, password string) error {
	/* 1. Check if row with a.ID exists; if it doesn't, return error
	 * 2. Sanitize password
	 * 3. Update password
	 */
	return nil
}

// Update `balance` column with `value`.
func SetBalance(conn *pgx.Conn, a Account, value float64) error {
	/* 1. Check if row with a.ID exists; if it doesn't, return error
	 * 2. Set balance
	 */
	return nil
}

// Check if provided credentials match.
func ValidateLogin(conn *pgx.Conn, login string, password string) error {
	/* 1. Check if login exists and password matches; if one of these fail,
	 * return error
	 * 2. On success, return nil
	 */
	return nil
}

// Uses TRANSACTION sql statement.
func DoTransaction(conn *pgx.Conn, a1 Account, a2 Account /* TODO: finish */) {

}
