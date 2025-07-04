package main

import "fmt"

func worker(ch <-chan int) {
	for val := range ch {
		fmt.Println("received:", val)
	}
	fmt.Println("channel closed, exit goroutine")
}

func main() {
	ch := make(chan int)

	go worker(ch)

	ch <- 1
	ch <- 2
	close(ch)
}
