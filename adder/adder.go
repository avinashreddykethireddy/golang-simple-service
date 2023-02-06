package adder

func Add(x, y int) int {
	return x + y
}

func CheckValues(x, y int) string {
	sum := x + y
	if sum < 10 {
		return "less"
	} else if sum == 0 {
		return "zero"
	} else {
		return "more"
	}
}
