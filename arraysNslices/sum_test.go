package arraysNslices

import "testing"

func TestSum(t *testing.T) {
	numbers := [5] int{1,2,3,4,5}

	got := Sum(numbers)
	want:= 15
	if want != got {
		t.Errorf("Wanted %d but Got %d was Given, %v", want, got, numbers)
	}
}
