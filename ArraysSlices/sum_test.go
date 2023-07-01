package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of sum of all collections ", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{13, 24, 31}
		numbers3 := []int{11, 24, 39}
		got := SumAll(numbers1, numbers2, numbers3)
		want := []int{6, 68, 74}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
