package datastructure

func FibonacciByFor(index int) int {
	if index == 0 {
		return 0
	} else {
		fibArr := make([]int, index+1)
		fibArr[0], fibArr[1] = 0, 1
		for i := 2; i <= index; i++ {
			fibArr[i] = fibArr[i-1] + fibArr[i-2]
		}
		return fibArr[index]
	}
}

func Fibonacci(index int) (fibN int) {
	if index == 0 {
		fibN = 0
	} else if index == 1 {
		fibN = 1
	} else {
		fibN = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return
}