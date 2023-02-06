package adder

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestCheckValues(t *testing.T) {
	ans1 := CheckValues(2, 2)
	expected1 := "less"
	ans2 := CheckValues(2, 9)
	expected2 := "more"

	if ans1 != expected1 {
		t.Errorf("expected '%s' but got '%s'", expected1, ans1)
	}

	if ans2 != expected2 {
		t.Errorf("expected '%s' but got '%s'", expected2, ans2)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
