package sqlx

import (
	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

type Pet struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Category string `db:"category"`
}

func SelectAllPets() ([]shelter.Pet, error) {
	query := `SELECT pet.id as id, pet.name as name, category.name as category FROM pet INNER JOIN category ON pet.category = category.id;`

	db, err := connect.SQLx()
	if err != nil {
		return nil, err
	}

	pets := make([]Pet, 0)
	err = db.Select(&pets, query)
	if err != nil {
		return nil, err
	}

	shelterPets := make([]shelter.Pet, len(pets))
	for i, pet := range pets {
		shelterPets[i] = shelter.Pet{
			Id:       pet.Id,
			Name:     pet.Name,
			Category: pet.Category,
		}
	}

	return shelterPets, nil
}

func SelectOnePet(id int) (shelter.Pet, error) {
	query := `SELECT pet.id as id, pet.name as name, category.name as category FROM pet INNER JOIN category ON pet.category = category.id WHERE pet.id = $1;`

	db, err := connect.SQLx()
	if err != nil {
		return shelter.Pet{}, err
	}

	pet := Pet{}
	err = db.Get(&pet, query, id)
	if err != nil {
		return shelter.Pet{}, err
	}

	return shelter.Pet{
		Id:       pet.Id,
		Name:     pet.Name,
		Category: pet.Category,
	}, nil
}
