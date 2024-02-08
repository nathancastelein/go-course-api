package sqlboiler

import (
	"github.com/nathancastelein/go-course-api/shelter"
)

func SelectAllPets() ([]shelter.Pet, error) {
	/*
		// Uncomment when you start!

		ctx := context.Background()
		db, err := connect.SQL()
		if err != nil {
			return nil, err
		}
	*/

	// List all pets

	shelterPets := make([]shelter.Pet, 0)

	// Cast all pets to shelter.Pet
	// Tips for Category: https://github.com/volatiletech/sqlboiler?tab=readme-ov-file#relationships

	return shelterPets, nil
}

func SelectOnePet(id int) (*shelter.Pet, error) {
	/*
		// Uncomment when you start!

		ctx := context.Background()
		db, err := connect.SQL()
		if err != nil {
			return nil, err
		}
	*/

	// Find one pet
	// Tips: https://github.com/volatiletech/sqlboiler?tab=readme-ov-file#find

	return &shelter.Pet{}, nil
}
