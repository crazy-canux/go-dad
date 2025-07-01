package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("end")
	panic("test panic")
	fmt.Println("this will not be printed")
}
