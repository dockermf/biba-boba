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
	Balance int64
}

// ERRORS
var (
	ErrInvalidInput    = errors.New("Username or password is empty")
	ErrUsernameTaken   = errors.New("Username already taken")
	ErrAccountNotFound = errors.New("Account not found")
)

// Insert new row into `Account` table.
func CreateAccount(ctx context.Context, db *pgxpool.Pool, accountType AccountType, username, password string) error {
	if username == "" || password == "" {
		return ErrInvalidInput
	}
	// TODO: sanitize password|username before doing SQL

	typeName, ok := AccountType_Name[accountType]
	if !ok {
		return fmt.Errorf("Invalid account type: %v", accountType)
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

// Retrieve account from `Account` table by ID.
func GetAccount(ctx context.Context, db *pgxpool.Pool, ID uint64) (*Account, error) {
	if ID == 0 {
		return nil, fmt.Errorf("Invalid account ID: %d", ID)
	}

	var acc Account
	var accType string

	err := db.QueryRow(
		ctx,
		"SELECT account_id, account_type, username, password, balance FROM Account WHERE account_id = $1",
		ID,
	).Scan(&acc.ID, &accType, &acc.Username, &acc.Password, &acc.Balance)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAccountNotFound
		} else {
			return nil, fmt.Errorf("Failed to get account: %w", err)
		}
	}

	if value, ok := AccountType_Value[accType]; ok {
		acc.AccountType = value
	} else {
		return nil, fmt.Errorf("Unknown account type from database: %s", accType)
	}

	return &acc, nil
}

// Delete row from `Account` table.
func RemoveAccount(ctx context.Context, db *pgxpool.Pool, ID uint64) error {
	if ID == 0 {
		return fmt.Errorf("Invalid account ID: %d", ID)
	}

	var accountExists uint64
	err := db.QueryRow(
		ctx,
		"SELECT account_id FROM Account WHERE account_id = $1",
		ID,
	).Scan(&accountExists)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// target account doesn't exist
			return ErrAccountNotFound
		} else {
			// random bullshit happened here, not recovering
			return fmt.Errorf("Expected pgx.ErrNoRows or nil, got: %w", err)
		}
	}
	// account exists, disintegrate that bitch
	_, err = db.Exec(
		ctx,
		"DELETE FROM Account WHERE account_id = $1",
		ID,
	)
	if err != nil {
		return fmt.Errorf("Failed to delete account: %w", err)
	}

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
