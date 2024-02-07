package sqlboiler

import (
	"context"

	"github.com/nathancastelein/go-course-api/database/sqlboiler/models"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func UpdatePetName(id int, newName string) error {
	ctx := context.Background()
	db, err := connect.SQL()
	if err != nil {
		return err
	}

	databasePet, err := models.FindPet(ctx, db, id)
	if err != nil {
		return err
	}

	databasePet.Name = newName
	_, err = databasePet.Update(ctx, db, boil.Infer())
	return err
}
