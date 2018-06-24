package sort

func Bubble(seq []int64) []int64 {
	for i := 0; i < len(seq)- 1; i++ {
		for j := 0; j < len(seq) - 1 -i; j++ {
			if seq[j] > seq[j+1] {
				seq[j], seq[j+1] = seq[j+1], seq[j]
            }
        }
    }
    return seq
}