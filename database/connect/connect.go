package connect

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func getPostgreSQLDSN() string {
	if pgdsn := os.Getenv("PG_DSN"); pgdsn != "" {
		return pgdsn
	}
	return "postgres://gocourse:test@localhost:5432/gocourse?sslmode=disable"
}

func SQL() (*sql.DB, error) {
	return sql.Open("postgres", getPostgreSQLDSN())
}
