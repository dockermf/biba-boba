package dbapi

import "errors"

// Account type, starts from 0. Can't be negative.
type AccountType uint32

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

// Errors
var (
	ErrInvalidID	      = errors.New("User ID is invalid")
	ErrInvalidAccountType = errors.New("Account type is invalid")
	ErrEmptyInput         = errors.New("User input is empty")
	ErrBadInput           = errors.New("User input has forbidden characters")
	ErrUsernameTaken      = errors.New("Username already taken")
	ErrAccountNotFound    = errors.New("Account not found")
)
