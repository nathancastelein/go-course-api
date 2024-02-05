package connect

import (
	"database/sql"
	"errors"
)

func SQL() (*sql.DB, error) {
	// Connection string: postgres://gocourse:test@localhost:5432/gocourse?sslmode=disable
	// Driver name: postgres
	// Documentation: https://pkg.go.dev/database/sql#Open
	return nil, errors.New("not implemented")
}
