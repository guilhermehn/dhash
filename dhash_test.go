package dhash

import (
	"testing"
)

func TestDhash(t *testing.T) {
	originalHash, err := Dhash("original.png", 8)

	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}

	modifiedHash, err := Dhash("modified.png", 8)

	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}

	if originalHash != modifiedHash {
		t.Errorf("Test failed - Original: \"%s\" / Modified: \"%s\"", originalHash, modifiedHash)
	}
}
