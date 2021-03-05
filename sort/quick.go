package sort

import "math/rand"

// QuickSort ...
func QuickSort(seq []int64) []int64 {
	if len(seq) <= 1 {
		return seq
	}
	pivot := seq[rand.Intn(len(seq))]
	greater := make([]int64, 0, len(seq))
	lesser := make([]int64, 0, len(seq))
	middle := make([]int64, 0, len(seq))

	for _, item := range seq {
		switch {
		case item < pivot:
			lesser = append(lesser, item)
		case item == pivot:
			middle = append(middle, item)
		case item > pivot:
			greater = append(greater, item)
		}
	}

	greater = QuickSort(greater)
	lesser = QuickSort(lesser)

	lesser = append(lesser, middle...)
	lesser = append(lesser, greater...)

	return lesser
}
