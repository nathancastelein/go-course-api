package sqlboiler

import (
	"context"

	"github.com/nathancastelein/go-course-api/database/sqlboiler/models"
	"github.com/nathancastelein/go-course-api/solution/database/connect"
)

func DeleteOnePet(id int) error {
	ctx := context.Background()
	db, err := connect.SQL()
	if err != nil {
		return err
	}

	databasePet, err := models.FindPet(ctx, db, id)
	if err != nil {
		return err
	}

	_, err = databasePet.Delete(ctx, db)
	return err
}
