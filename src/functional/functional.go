package functional

func chop(elem int, arr []int) int {
	// Initialize index and sliceLength
	index := 0
	sliceLength := len(arr)
	for sliceLength != 0 {
		if sliceLength == 1 {
			// Check if the only element in the slice is the one we're looking for
			if arr[0] == elem {
				return index
			}
			break
		}
		index, arr = chopSlice(elem, arr, index)
		sliceLength = len(arr)
	}
	// The element was not found
	return -1
}

func chopSlice(elem int, arr []int, index int) (newIndex int, newSlice []int) {
	head := arr[0 : len(arr)/2]
	tail := arr[len(arr)/2:]
	// Knowing that the slice is sorted, find which half should contain the element
	if tail[0] <= elem {
		return index + len(head), tail
	}
	return index, head
}
