package structural

type chopInput struct {
	arr []int
}

func (i *chopInput) chop(elem int) int {
	// Initialize the start and end marks
	start := 0
	end := len(*&i.arr)
	for start < end {
		// Find the middle index
		middle := (start + end) / 2
		if i.arr[middle] == elem {
			// The item in the middle of the slice is the element we're searching
			return middle
		} else if elem > i.arr[middle] {
			// The element should be in the second half
			start = middle + 1
		} else {
			// The element should be in the first half
			end = middle
		}
	}

	// The element was not found
	return -1
}
