package integers

import (
	"fmt"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	returnedValue := AddNumbers(1, 2)
	expectedValue := 3
	if returnedValue != expectedValue {
		t.Errorf("got '%d', wanted '%d'", returnedValue, expectedValue)
	}
}

func ExampleAddNumbers() {
	sum := AddNumbers(1, 2)
	fmt.Println(sum)
	// Output: 3
}
