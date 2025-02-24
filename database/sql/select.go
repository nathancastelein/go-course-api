package sql

import (
	"errors"
	"log/slog"

	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

func SimpleQuery() (bool, error) {
	query := `SELECT 1 = 1;`

	db, err := connect.SQL()
	if err != nil {
		return false, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		var result bool
		if err := rows.Scan(&result); err != nil {
			return false, err
		}

		return result, nil
	}

	return false, errors.New("no data")
}

func SelectAllPetNames() ([]string, error) {
	query := `SELECT pet.name FROM pet;`
	slog.Debug("executing query", slog.String("query", query))

	return nil, errors.New("not implemented")
}

func SelectAllPets() ([]shelter.Pet, error) {
	query := `SELECT pet.id, pet.name, category.name FROM pet INNER JOIN category ON pet.category = category.id;`

	db, err := connect.SQL()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	pets := make([]shelter.Pet, 0)
	for rows.Next() {
		// TODO: Instanciate a new empty pet
		// TODO: Scan in the proper fields
		// TODO: Append the pet in the result
	}

	return pets, nil
}
