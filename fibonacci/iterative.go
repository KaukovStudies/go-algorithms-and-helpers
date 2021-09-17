package fibonacci

func Iterative(n int) int {
	if n < 1 {
		return 1
	}

	calculated := []int{1, 1}

	for len(calculated) < n {
		currentLength := len(calculated)

		calculated = append(calculated, calculated[currentLength-1]+calculated[currentLength-2])
	}

	return calculated[n-1]
}
