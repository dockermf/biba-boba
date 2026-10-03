// Package dbapi is used to talk with Postgres
package dbapi // import "github.com/dockermf/biba-boba/include/dbapi"

import (
    "context"
    "errors"
    "fmt"

    "github.com/jackc/pgx/v5"
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
    AccountType_Name = map[AccountType]string {
        0: "Regular",
        1: "Admin",
    }
    AccountType_Value = map[string]AccountType {
        "Regular": 0,
        "Admin":   1,
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

// Insert new row into `Account` table.
func CreateAccount(conn *pgx.Conn, account_type AccountType, username string, password string) error {
    if username == "" || password == "" {
        return errors.New("Login or password empty")
    }

    if _, ok := AccountType_Name[account_type]; !ok {
        return fmt.Errorf("Account type %v is invalid", account_type)
    }

    var existingAccount string
    err := conn.QueryRow(
        context.Background(),
        "SELECT username FROM Account WHERE username = $1",
        username,
    ).Scan(&existingAccount)

    if err == nil {
        return fmt.Errorf("Name '%s' already taken", existingAccount)
    }

    // Weird line?
    if !errors.Is(err, pgx.ErrNoRows) {
        return fmt.Errorf("Expected error pgx.ErrNoRows, got: %v\n", err)
    }

    // TODO: hash password
    _, err = conn.Exec(
        context.Background(),
        "INSERT INTO Account (account_type, login, password) VALUES ($1, $2, $3)",
        account_type,
        username,
        password,
    )
    if err != nil {
        return fmt.Errorf("Failed to create new account: %v", err)
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
func DoTransaction(conn *pgx.Conn, a1 Account, a2 Account/* TODO: finish */) {

}
