package sort

// SelectionSort ...
func SelectionSort(seq []int64) []int64 {
	for i := 0; i < len(seq)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(seq); j++ {
			if seq[j] < seq[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			seq[i], seq[minIndex] = seq[minIndex], seq[i]
		}
	}
	return seq
}
