package dp

func fibonacci(n uint, memo []uint64) uint64 {
	if n == 1 || n == 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)

	return memo[n]
}

//FibonacciFast : return memoized fibonacci
func FibonacciRecursive(n uint) uint64 {
	var memo = make([]uint64, n+1)

	return fibonacci(n, memo)

}

func FibonacciIterative(n uint) uint64 {
	var memo = make([]uint64, n+1)

	memo[1], memo[2] = 1, 1

	var i uint = 3

	for ; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]

}

func FibonacciFast(n uint) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	var prev, curr uint64
	prev, curr = 1, 1
	var i uint = 3
	for ; i <= n; i++ {
		var sum uint64 = prev + curr
		prev, curr = curr, sum
	}
	return curr
}
