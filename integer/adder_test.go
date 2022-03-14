package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expecting := 4

	if sum != expecting{
		t.Errorf("Expecting '%d' but got '%d' ", expecting, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}
