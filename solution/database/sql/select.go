package sql

import (
	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

func SelectAllPetNames() ([]string, error) {
	query := `SELECT pet.name FROM pet;`

	db, err := connect.SQL()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	petNames := make([]string, 0)
	for rows.Next() {
		var petName string
		if err := rows.Scan(&petName); err != nil {
			return nil, err
		}

		petNames = append(petNames, petName)
	}

	return petNames, nil
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
		pet := shelter.Pet{}
		if err := rows.Scan(&pet.Id, &pet.Name, &pet.Category); err != nil {
			return nil, err
		}

		pets = append(pets, pet)
	}

	return pets, nil
}
