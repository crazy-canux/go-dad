package search

// LinearSearch ...
func LinearSearch(seq []int64, keyword int64) int {
	for index, value := range seq {
		if value == keyword {
			return index
		}
	}
	return -1
}
