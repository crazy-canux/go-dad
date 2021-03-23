package sort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	seq := []int64{23, 78, -25, 82, 0, 589}
	if res := BubbleSort(seq); res[0] != -25 || res[5] != 589 {
		fmt.Println(res)
		t.Error("Failed: sort failed.")
	} else {
		fmt.Println(res)
	}
}

func TestInsertion(t *testing.T) {
	seq := []int64{23, 78, -25, 82, 0, 589}
	if res := InsertionSort(seq); res[0] != -25 || res[5] != 589 {
		fmt.Println(res)
		t.Error("Failed: sort failed.")
	} else {
		fmt.Println(res)
		t.Log("sort succeed.")
	}
}

func TestQuickSort(t *testing.T) {
	seq := []int64{23, 78, -25, 82, 0, 589}
	if res := QuickSort(seq); res[0] != -25 || res[5] != 589 {
		fmt.Println(res)
		t.Error("Failed: sort failed.")
	} else {
		fmt.Println(res)
	}
}

func TestSelection(t *testing.T) {
	seq := []int64{23, 78, -25, 82, 0, 589}
	if res := SelectionSort(seq); res[0] != -25 && res[5] != 589 {
		fmt.Println(res)
		t.Error("Selection sort failed.")
	} else {
		fmt.Println(res)
	}
}

func TestShell(t *testing.T) {
	seq := []int64{23, 78, -25, 82, 0, 589}
	if res := ShellSort(seq); res[0] != -25 && res[5] != 589 {
		fmt.Print(res)
		t.Error("shell sort failed")
	} else {
		fmt.Print(res)
		t.Log("shell sort succeed")
	}
}
