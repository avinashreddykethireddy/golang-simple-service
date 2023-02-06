package adder

func Add(x, y int) int {
	return x + y
}

func CheckValues(x, y int) string {
	sum := x + y
	if sum > 10 {
		return "more"
	} else if sum == 0 {
		return "zero"
	} else {
		return "less"
	}
}

func SubtractAndSay(x, y int) string {
	sub := x - y
	if sub < 0 {
		return "less than zero"
	} else if sub == 0 {
		return "equal to zero"
	} else {
		return "more than zero"
	}
}
