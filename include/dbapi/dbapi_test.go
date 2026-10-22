package dbapi_test

import (
	"context"
	"testing"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
    . "github.com/dockermf/biba-boba/include/dbapi"
)

// TestInvalidID calls (Remove|Get)Account functions with invalid user IDs.
func TestInvalidID(t *testing.T) {
	// DATABASE_URL = postgres://DB_USERNAME:DB_PASSWORD@DB_HOST:DB_PORT/DB_NAME
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err = RemoveAccount(context.Background(), pool, 0)
	if err == nil {
		t.Errorf("RemoveAccount(context.Background(), pool, 0) = nil, want %v", ErrInvalidID)
	}

	acc, err := GetAccount(context.Background(), pool, 0)
	if acc != nil {
		t.Errorf("GetAccount(context.Background(), pool, 0) = %v, %v, want nil, %v", acc, err, ErrInvalidID)
	}

	defer pool.Close()
}
