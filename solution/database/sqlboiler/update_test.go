package sqlboiler

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	// Arrange
	// Act
	err := UpdatePetName(1, "Hoagie")

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}
}
