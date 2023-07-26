package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats default five times if given zero", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		assertExpectedValue(t, expected, repeated)
	})

	t.Run("repeats default 5 times if negative value", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"

		assertExpectedValue(t, expected, repeated)
	})

	t.Run("repeats given n times", func(t *testing.T) {
		times := 3
		repeated := Repeat("a", times)
		expected := "aaa"

		assertExpectedValue(t, expected, repeated)
	})
}

func assertExpectedValue(t *testing.T, expected string, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func ExampleRepeat() {
	repeatedCharacter := Repeat("a", 4)
	fmt.Println(repeatedCharacter)
	//Output: aaaa
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
