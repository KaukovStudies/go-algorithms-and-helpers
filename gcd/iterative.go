package gcd

func Iterative(a, b int) int {
	r := 0

	for (a % b) > 0 {
		r = a % b
		a = b
		b = r
	}

	return b
}
