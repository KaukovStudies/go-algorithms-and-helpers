package gcd

func Recursive(a, b int) int {
	if b == 0 {
		return a
	}

	return Recursive(b, a%b)
}
