package connect

import (
	"database/sql"
	"errors"
	"os"
)

func getPostgreSQLDSN() string {
	if pgdsn := os.Getenv("PG_DSN"); pgdsn != "" {
		return pgdsn
	}
	return "postgres://gocourse:test@localhost:5432/gocourse?sslmode=disable"
}

func SQL() (*sql.DB, error) {
	// Connection string: use getPostgresqlDSN function.
	// Driver name: postgres
	// Documentation: https://pkg.go.dev/database/sql#Open
	return nil, errors.New("not implemented")
}
