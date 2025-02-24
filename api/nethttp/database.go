package main

import (
	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/sqlboiler"
)

type petRepository struct{}

func (p petRepository) SelectAllPets() ([]shelter.Pet, error) {
	return sqlboiler.SelectAllPets()
}

func (p petRepository) SelectOnePet(id int) (*shelter.Pet, error) {
	return sqlboiler.SelectOnePet(id)
}

func (p petRepository) UpdatePetName(id int, newName string) error {
	return sqlboiler.UpdatePetName(id, newName)
}

func newPetRepository() shelter.PetRepository {
	return &petRepository{}
}
