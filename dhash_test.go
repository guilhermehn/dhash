package dhash

import (
	"fmt"
	"testing"
)

func TestDhash(t *testing.T) {
	expectedHash := "4c8e3366c275650f"
	resultHash, _ := Dhash("test.jpg", 8)

	fmt.Println(Dhash("test.jpg", 8))
	fmt.Println(Dhash("test.jpg", 8))
	fmt.Println(Dhash("test.jpg", 8))
	fmt.Println(Dhash("test.jpg", 8))

	if resultHash != expectedHash {
		t.Errorf("Test failed - Expected \"%s\" got \"%s\"", expectedHash, resultHash)
	}
}
