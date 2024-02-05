package connect

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getPostgreSQLDSN() string {
	if pgdsn := os.Getenv("PG_DSN"); pgdsn != "" {
		fmt.Println("USING PGDSN")
		return pgdsn
	}
	return "postgres://gocourse:test@localhost:5432/gocourse?sslmode=disable"
}

func SQL() (*sql.DB, error) {

	db, err := sql.Open("postgres", getPostgreSQLDSN())
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

func SQLx() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", getPostgreSQLDSN())
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
