package sql

import (
	"errors"

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

	for rows.Next() {
		var result bool
		if err := rows.Scan(&result); err != nil {
			return false, err
		}

		return result, nil
	}

	return false, errors.New("no data")
}

func SelectAllPetNames() ([]string, error) {
	return nil, errors.New("not implemented")
}

func SelectAllPets() ([]shelter.Pet, error) {
	query := `` // Tips: your query will require a JOIN

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
