package recursive

func chop(elem int, arr []int) int {
	if len(arr) == 0 {
		// Empty slice
		return -1
	}
	return rChop(elem, arr, 0)
}

func rChop(elem int, arr []int, index int) int {
	if len(arr) == 1 {
		// Slice of size 1
		if arr[0] == elem {
			return index
		}
		return -1
	}

	// Split the slice at the middle
	head := arr[0 : len(arr)/2]
	tail := arr[len(arr)/2:]
	// Knowing that the slice is sorted, find which half should contain the element
	if tail[0] <= elem {
		index = index + len(head)
		return rChop(elem, tail, index)
	}
	return rChop(elem, head, index)
}
