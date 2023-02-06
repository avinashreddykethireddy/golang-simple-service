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

func TestCheckValuesLess(t *testing.T) {
	ans1 := CheckValues(2, 2)
	expected1 := "less"
	if ans1 != expected1 {
		t.Errorf("expected '%s' but got '%s'", expected1, ans1)
	}
}

func TestCheckValuesMore(t *testing.T) {
	ans2 := CheckValues(2, 9)
	expected2 := "more"
	if ans2 != expected2 {
		t.Errorf("expected '%s' but got '%s'", expected2, ans2)
	}
}

func TestEqual(t *testing.T) {
	ans3 := CheckValues(2, -2)
	expected3 := "zero"
	if ans3 != expected3 {
		t.Errorf("expected '%s' but got '%s'", expected3, ans3)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
