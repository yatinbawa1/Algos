package fibonacci

func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a := 0
	b := 1
	for i := 2; i < n; i++ {
		a, b = b, b+a
	}

	// is there a way to make this better
	// Can I have a single variable go
	return a + b
}
