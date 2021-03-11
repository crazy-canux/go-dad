package sort

// InsertionSort ...
func InsertionSort(seq []int64) []int64 {
	for i := 1; i < len(seq); i++ {
		insertIndex := i
		insertValue := seq[i]
		for ; insertIndex > 0 && insertValue <= seq[insertIndex-1]; insertIndex-- {
			seq[insertIndex] = seq[insertIndex-1]
		}
		seq[insertIndex] = insertValue
	}
	return seq
}
