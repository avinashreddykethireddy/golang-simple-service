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
	cases := []struct {
		x        int
		y        int
		expected string
	}{
		{2, 5, "less"},
		{2, 10, "more"},
		{2, -2, "zero"},
		{2, -7, "less"},
	}
	for _, c := range cases {
		t.Run("testing", func(t *testing.T) {
			got := CheckValues(c.x, c.y)
			if got != c.expected {
				t.Fatalf("Wanted %v , but got %v", c.expected, got)
			}
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

//func TestSubtractAndSay(t *testing.T) {
//	ans := SubtractAndSay(2, 3)
//	expected := "less than zero"
//
//	if ans != expected {
//		t.Errorf("expected '%s' but got '%s'", expected, ans)
//	}
//}
