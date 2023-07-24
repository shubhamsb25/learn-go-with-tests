package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expected %d found %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(4, 5)
	fmt.Println(sum)
	// Output: 9
}
