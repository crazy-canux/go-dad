package datastructure

func FibonacciByFor(index int) int {
	if index <= 1 {
		return index
	}
	FibArr := make([]int, index+1)
	FibArr[0], FibArr[1] = 0, 1
	for i := 2; i < index+1; i++ {
		FibArr[i] = FibArr[i-1] + FibArr[i-2]
	}
	return FibArr[index]
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
