package connect

import "testing"

func TestConnectSQL(t *testing.T) {
	// Arrange
	// Act
	db, err := SQL()

	// Assert
	if err != nil {
		t.Fatalf("err = '%s'; want nil", err)
	}

	if db == nil {
		t.Fatalf("db = nil; want not nil")
	}
}
