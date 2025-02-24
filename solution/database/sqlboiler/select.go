package sqlboiler

import (
	"context"

	"github.com/nathancastelein/go-course-api/database/sqlboiler/models"
	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

func SelectAllPets() ([]shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.SQL()
	if err != nil {
		return nil, err
	}

	pets, err := models.Pets().All(ctx, db)
	if err != nil {
		return nil, err
	}

	shelterPets := make([]shelter.Pet, len(pets))
	for i, pet := range pets {
		category, err := pet.PetCategory().One(ctx, db)
		if err != nil {
			return nil, err
		}
		shelterPets[i] = shelter.Pet{
			Id:       pet.ID,
			Name:     pet.Name,
			Category: category.Name,
		}
	}

	return shelterPets, nil
}

func SelectOnePet(id int) (*shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.SQL()
	if err != nil {
		return nil, err
	}

	pet, err := models.FindPet(ctx, db, id)
	if err != nil {
		return nil, err
	}

	category, err := pet.PetCategory().One(ctx, db)
	if err != nil {
		return nil, err
	}
	return &shelter.Pet{
		Id:       pet.ID,
		Name:     pet.Name,
		Category: category.Name,
	}, nil
}

func SelectByCategories(categories ...string) ([]shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.SQL()
	if err != nil {
		return nil, err
	}

	databaseCategories, err := models.Categories(models.CategoryWhere.Name.IN(categories)).All(ctx, db)
	if err != nil {
		return nil, err
	}

	categoriesIds := make([]int, len(databaseCategories))
	for i, category := range databaseCategories {
		categoriesIds[i] = category.ID
	}

	pets, err := models.Pets(models.PetWhere.Category.IN(categoriesIds)).All(ctx, db)
	if err != nil {
		return nil, err
	}

	shelterPets := make([]shelter.Pet, len(pets))
	for i, pet := range pets {
		category, err := pet.PetCategory().One(ctx, db)
		if err != nil {
			return nil, err
		}
		shelterPets[i] = shelter.Pet{
			Id:       pet.ID,
			Name:     pet.Name,
			Category: category.Name,
		}
	}

	return shelterPets, nil
}
