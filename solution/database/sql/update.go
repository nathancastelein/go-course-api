package sql

import (
	"errors"

	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

func UpdatePetName() error {
	query := `UPDATE pet SET name = $1 WHERE id = $2`

	db, err := connect.SQL()
	if err != nil {
		return err
	}

	result, err := db.Exec(query, "Bernard", 4)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("update failed, no rows affected")
	}

	return nil
}
