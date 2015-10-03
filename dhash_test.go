package dhash

import (
	"testing"
)

func TestDhash(t *testing.T) {
	size := 8

	originalHash, err := Dhash("original.png", size)

	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}

	modifiedHash, err := Dhash("modified.png", size)

	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}

	t.Logf("\nOriginal file hash: %s\nModified file hash: %s", originalHash, modifiedHash)

	if originalHash != modifiedHash {
		t.Errorf("Test failed - Original: \"%s\" / Modified: \"%s\"", originalHash, modifiedHash)
	}
}
