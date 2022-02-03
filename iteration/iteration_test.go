package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeat(t *testing.T) {
	returnedValue := Repeat("a")
	expectedValue := "aaaaa"

	if returnedValue != expectedValue {
		t.Errorf("got %q, wanted %q", returnedValue, expectedValue)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}
