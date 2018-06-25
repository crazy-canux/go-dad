package sort

import (
	"testing"
	"fmt"
)

func TestSelection(t *testing.T) {
	seq := []int64{23, 78, -25, 82, 0, 589}
	if res := Selection(seq); res[0] != -25 && res[5] != 589 {
		fmt.Println(res)
		t.Error("Selection sort failed.")
	} else {
		fmt.Println(res)
	}
}