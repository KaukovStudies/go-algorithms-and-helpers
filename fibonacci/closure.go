package fibonacci

func Closure() func() int {
	prev := 1
	next := 1

	return func() int {
		res := prev

		prev, next = next, (prev + next)

		return res
	}
}
