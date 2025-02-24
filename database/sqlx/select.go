package sqlx

import (
	"log/slog"

	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

func SelectAllPets() ([]shelter.Pet, error) {
	query := `SELECT pet.id as id, pet.name as name, category.name as category FROM pet INNER JOIN category ON pet.category = category.id;`

	db, err := connect.SQLx()
	if err != nil {
		return nil, err
	}

	slog.Debug("executing query", slog.String("query", query), slog.Any("db", db))

	return nil, nil
}

func SelectOnePet(id int) (*shelter.Pet, error) {
	query := `SELECT pet.id as id, pet.name as name, category.name as category FROM pet INNER JOIN category ON pet.category = category.id WHERE pet.id = $1;`

	db, err := connect.SQLx()
	if err != nil {
		return nil, err
	}

	slog.Debug("executing query", slog.String("query", query), slog.Any("db", db))

	return nil, nil
}
