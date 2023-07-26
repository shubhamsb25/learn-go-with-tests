package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of given collection", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum of given multiple collections", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3}

		got := SumAll(numbers1, numbers2)
		want := []int{15, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
