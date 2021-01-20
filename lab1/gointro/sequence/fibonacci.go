package sequence

// Task: Fibonacci numbers
//
// fibonacci(n) returns the n-th Fibonacci number, and is defined by the
// recurrence relation F_n = F_n-1 + F_n-2, with seed values F_0=0 and F_1=1.

func sum(a, b uint) uint {
	return a + b
}

func fibonacci(n uint) uint {
	var i uint
	fib := make([]uint, n+1)

	if n == 0 || n == 1 {
		return n
	}
	fib[0], fib[1] = 0, 1

	for i = 1; i < n; i++ {
		fib[i+1] = sum(fib[i], fib[i-1])
	}

	return fib[n]
}
