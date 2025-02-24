package sql

import (
	"errors"
	"log/slog"
)

func UpdatePetName(id int, newName string) error {
	query := `UPDATE pet SET name = $1 WHERE id = $2`
	slog.Debug("executing query", slog.String("query", query))

	return errors.New("not implemented")
}
