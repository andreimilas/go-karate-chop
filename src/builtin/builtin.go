package builtin

import "sort"

func chop(elem int, arr []int) int {
	// The SearchInts function performs a binary search for an item in a sorted slice
	// If the element is not found, an index where the element should be inserted is returned
	index := sort.SearchInts(arr, elem)

	if index != len(arr) && arr[index] == elem {
		// If we haven't passed the end of the slice AND the element was found, return the index
		return index
	}

	return -1
}
