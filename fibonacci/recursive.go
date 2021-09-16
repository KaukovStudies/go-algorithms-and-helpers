package fibonacci

func Recursive(n int) int {
	if n <= 2 {
		return 1
	}

	return Recursive(n-1) + Recursive(n-2)
}
