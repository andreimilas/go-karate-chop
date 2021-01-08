package iterative

func chop(elem int, arr []int) int {
	// Initialize start and end marks
	start := 0
	end := len(arr)
	for start < end {
		middle := (start + end) / 2
		if elem == arr[middle] {
			return middle
		} else if elem > arr[middle] {
			start = middle + 1
		} else {
			end = middle
		}
	}
	// The element was not found
	return -1
}
