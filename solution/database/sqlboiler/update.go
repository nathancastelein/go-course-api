package sqlboiler

import (
	"context"

	"github.com/nathancastelein/go-course-api/database/sqlboiler/models"
	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func UpdatePet(pet *shelter.Pet) error {
	ctx := context.Background()
	db, err := connect.SQL()
	if err != nil {
		return err
	}

	databasePet, err := models.FindPet(ctx, db, pet.Id)
	if err != nil {
		return err
	}

	databasePet.Name = pet.Name
	_, err = databasePet.Update(ctx, db, boil.Infer())
	return err
}
