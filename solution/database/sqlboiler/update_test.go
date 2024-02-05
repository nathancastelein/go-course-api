package sqlboiler

import (
	"testing"

	"github.com/nathancastelein/go-course-api/shelter"
)

func TestUpdate(t *testing.T) {
	// Arrange
	// Act
	err := UpdatePet(&shelter.Pet{Id: 1, Name: "Hoagie", Category: "Dog"})

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}
}
