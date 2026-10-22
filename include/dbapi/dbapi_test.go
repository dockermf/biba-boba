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
		t.Errorf("RemoveAccount(context.Background(), pool, 0) = nil, want %v", ErrInvalidID)
	}

	acc, err := GetAccount(context.Background(), pool, 0)
	if acc != nil {
		t.Errorf("GetAccount(context.Background(), pool, 0) = %v, %v, want nil, %v", acc, err, ErrInvalidID)
	}

	defer pool.Close()
}
