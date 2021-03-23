package sort

// ShellSort ...
func ShellSort(seq []int64) []int64 {
	for gap := int(len(seq) / 2); gap > 0; gap /= 2 {
		for i := gap; i < len(seq); i++ {
			for j := i; j >= gap && seq[j-gap] > seq[j]; j -= gap {
				seq[j], seq[j-gap] = seq[j-gap], seq[j]
			}
		}
	}
	return seq
}
