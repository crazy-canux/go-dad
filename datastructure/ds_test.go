package datastructure

import (
	"testing"
)

func TestFibonacciByFor(t *testing.T) {
	if res := FibonacciByFor(0); res != 0 {
		t.Error("Error: get fib 0 failed.")
	} else {
		t.Log("Pass: fib 0: ", res)
	}

	if res := FibonacciByFor(1); res != 1 {
		t.Error("Error: get fib 1 failed.")
	} else {
		t.Log("Pass: fib 1: ", res)
	}

	if res := FibonacciByFor(2); res != 1 {
		t.Error("Error: get fib 2 failed.")
	} else {
		t.Log("Pass: fib 2: ", res)
	}

	if res := FibonacciByFor(3); res != 2 {
		t.Error("Error: get fib 3 failed.")
	} else {
		t.Log("Pass: fib 3: ", res)
	}
}

func TestFibonacci(t *testing.T) {
	if res := Fibonacci(0); res != 0 {
		t.Error("Error: get fib 0 failed.")
	} else {
		t.Log("Pass: fib 0: ", res)
	}

	if res := Fibonacci(1); res != 1 {
		t.Error("Error: get fib 1 failed.")
	} else {
		t.Log("Pass: fib 1: ", res)
	}

	if res := Fibonacci(2); res != 1 {
		t.Error("Error: get fib 2 failed.")
	} else {
		t.Log("Pass: fib 2: ", res)
	}

	if res := Fibonacci(3); res != 2 {
		t.Error("Error: get fib 3 failed.")
	} else {
		t.Log("Pass: fib 3: ", res)
	}
}

func TestLinkedListReverse(t *testing.T) {
	// Create a linked list: 1 -> 2 -> 3
	list := &LinkedList{
		Head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	reversed := Reverse(list.Head)
	if reversed.Val != 3 || reversed.Next.Val != 2 || reversed.Next.Next.Val != 1 {
		t.Error("Error: LinkedList reverse failed.")
	} else {
		t.Log("Pass: LinkedList reverse: ", reversed)
	}
}
