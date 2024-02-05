package sqlboiler

import "testing"

func TestSelectAllPets(t *testing.T) {
	// Arrange

	// Act
	pets, err := SelectAllPets()

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}

	if pets == nil {
		t.Fatal("pets is nil; want not nil")
	}

	if len(pets) == 0 {
		t.Fatal("pets is empty; want not empty")
	}
}

func TestSelectOnePet(t *testing.T) {
	// Arrange

	// Act
	pet, err := SelectOnePet(1)

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}

	if pet == nil {
		t.Fatal("pet is nil; want not nil")
	}

	if pet.Id != 1 {
		t.Fatalf("pet id is %d; want 1", pet.Id)
	}
}
