package search

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	seq1 := []int64{23, 78, -25, 82, 0, 589}
	if res := BinarySearch(seq1, -25); res == -1 {
		t.Error("Failed: search -25 failed.")
	} else {
		fmt.Println(res)
	}
}

func TestSequential(t *testing.T) {
	seq1 := []int64{23, 78, -25, 82, 0, 589}
	if res := SequentialSearch(seq1, -25); res == -1 {
		t.Error("Failed: search -25 failed.")
	} else {
		fmt.Println(res)
	}
}
