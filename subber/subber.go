package subber

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
