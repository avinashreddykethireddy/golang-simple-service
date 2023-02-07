package subber

import (
	"testing"
)

func TestSubtractAndSay(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected string
	}{
		{2, 5, "less than zero"},
		{2, 1, "more than zero"},
		//{2, -2, "zero"},
		//{2, -7, "less"},
	}
	for _, c := range cases {
		t.Run("testing", func(t *testing.T) {
			got := SubtractAndSay(c.x, c.y)
			if got != c.expected {
				t.Fatalf("Wanted %v , but got %v", c.expected, got)
			}
		})
	}
}
