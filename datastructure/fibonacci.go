package datastructure

func FibonacciByFor(index int) int {
	fibArr := make([]int, index)
	fibArr[0], fibArr[1]  = 1, 1

	for i := 2; i < index; i++ {
		fibArr[i] = fibArr[i-1] + fibArr[i-2]
	}
	return fibArr[index]
}

func Fibonacci(index int) (fibN int) {
	if index <= 2 {
		fibN = 1
	} else {
		fibN = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return
}
