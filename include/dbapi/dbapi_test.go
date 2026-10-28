package dbapi_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
    . "github.com/dockermf/biba-boba/include/dbapi"
)

// TestInvalidID calls (Remove|Get)Account functions with invalid user IDs.
func TestInvalidID(t *testing.T) {
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

	err = RemoveAccount(context.Background(), pool, 0)
	if err == nil {
		t.Errorf("Remove id 0: nil, want [%v]", ErrInvalidID)
	}

	acc, err := GetAccount(context.Background(), pool, 0)
	if err == nil {
		t.Errorf("Get account id 0: [%v], [%v], want nil, [%v]", acc, err, ErrInvalidID)
	}

	defer pool.Close()
}

func TestCreateAccount(t *testing.T) {
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

	// Test invalid account type case
	err = CreateAccount(context.Background(), pool, -1, "user", "password")
	if err != ErrInvalidAccountType {
		t.Log(err)
		t.Errorf("Account type -1: [%v], want [%v]", err, ErrInvalidAccountType)
	}

	// Test empty username
	err = CreateAccount(context.Background(), pool, AccountType_Regular, "", "something")
	if err != ErrEmptyInput {
		t.Errorf("Empty username: [%v], want [%v]", err, ErrEmptyInput)
	}

	// Test blank username (2 tabs)
	err = CreateAccount(context.Background(), pool, AccountType_Regular, "\t\t", "something")
	if err != ErrEmptyInput {
		t.Errorf("Blank username: [%v], want [%v]", err, ErrEmptyInput)
	}

	// Test empty password
	err = CreateAccount(context.Background(), pool, AccountType_Regular, "user", "")
	if err != ErrEmptyInput {
		t.Errorf("Empty password: [%v], want [%v]", err, ErrEmptyInput)
	}

	// Test blank password (a newline)
	err = CreateAccount(context.Background(), pool, AccountType_Regular, "user", "\n")
	if err != ErrEmptyInput {
		t.Errorf("Blank password: [%v], want [%v]", err, ErrEmptyInput)
	}

	defer pool.Close()
}
