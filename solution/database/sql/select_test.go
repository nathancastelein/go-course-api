package sql

import "testing"

func TestSelectAllPetNames(t *testing.T) {
	// Arrange

	// Act
	petNames, err := SelectAllPetNames()

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}

	if petNames == nil {
		t.Fatal("petNames is nil; want not nil")
	}

	if len(petNames) == 0 {
		t.Fatal("petNames is empty; want not empty")
	}
}

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
