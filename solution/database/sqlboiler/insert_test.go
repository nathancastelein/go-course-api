package sqlboiler

import (
	"testing"

	"github.com/nathancastelein/go-course-api/shelter"
)

func TestInsert(t *testing.T) {
	// Arrange

	// Act
	pet, err := InsertNewPet(&shelter.Pet{
		Name:     "Medor",
		Category: "Dog",
	})

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}

	if pet == nil {
		t.Fatal("pet is nil; want not nil")
	}

	if pet.Id == 0 {
		t.Fatalf("pet id is %d; want not zero", pet.Id)
	}
}
