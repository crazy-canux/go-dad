package search

// BinarySearch ...
func BinarySearch(orderSeq []int64, keyword int64) int {
	minIdx := 0
	maxIdx := len(orderSeq) - 1
	for minIdx <= maxIdx {
		midIdx := (minIdx + maxIdx) >> 1
		if keyword > orderSeq[midIdx] {
			minIdx = midIdx + 1
		} else if keyword < orderSeq[midIdx] {
			maxIdx = midIdx - 1
		} else {
			return midIdx
		}
	}
	return -1
}
