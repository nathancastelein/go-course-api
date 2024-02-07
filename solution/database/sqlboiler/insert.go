package sqlboiler

import (
	"context"

	"github.com/nathancastelein/go-course-api/database/sqlboiler/models"
	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InsertNewPet(pet shelter.Pet) (shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.SQL()
	if err != nil {
		return shelter.Pet{}, err
	}

	category, err := models.Categories(models.CategoryWhere.Name.EQ(pet.Category)).One(ctx, db)
	if err != nil {
		return shelter.Pet{}, err
	}

	databasePet := models.Pet{
		Name:     pet.Name,
		Category: null.NewInt(category.ID, true),
	}

	if err := databasePet.Insert(ctx, db, boil.Infer()); err != nil {
		return shelter.Pet{}, err
	}

	return shelter.Pet{
		Id:       databasePet.ID,
		Name:     databasePet.Name,
		Category: category.Name,
	}, nil
}
