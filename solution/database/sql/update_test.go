package sql

import "testing"

func TestUpdate(t *testing.T) {
	// Arrange
	// Act
	err := UpdatePetName(4, "Bernard")

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}
}
