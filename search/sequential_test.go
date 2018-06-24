package search

import (
	"testing"
	"fmt"
)

func TestSequential(t *testing.T) {
	seq1 := []int64{23, 78, -25, 82, 0, 589}
	if res := Sequential(seq1, -25); res == -1 {
		t.Error("Failed: search -25 failed.")
	} else {
		fmt.Println(res)
	}
}