package sqlx

import (
	"fmt"

	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

func SelectAllPets() ([]shelter.Pet, error) {
	query := ``

	db, err := connect.SQLx()
	if err != nil {
		return nil, err
	}

	fmt.Println("I want to execute a query on db", query, db) // Remove me
	return nil, nil
}

func SelectOnePet(id int) (*shelter.Pet, error) {
	query := ``

	db, err := connect.SQLx()
	if err != nil {
		return nil, err
	}

	fmt.Println("I want to execute a query on db", query, db) // Remove me
	return nil, nil
}
