package arraysNslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of specified numbers", func(t *testing.T) {
		numbers := [5] int{1,2,3,4,5}

		got := Sum(numbers)
		want:= 15
		if want != got {
			t.Errorf("Wanted %d but Got %d was Given, %v", want, got, numbers)
		}
	})

	t.Run("Collection for unspecified number of items", func(t *testing.T) {
		numbers := []int{1,2,3}

		got := SumUnspecified(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d but was expecting %d and was given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %v got %v", want, got)
	}
}
