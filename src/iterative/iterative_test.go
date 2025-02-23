package iterative

import "testing"

func TestChop(t *testing.T) {
	var tests = []struct {
		elem          int
		arr           []int
		expectedIndex int
	}{
		{3, []int{}, -1},
		{3, []int{1}, -1},
		{1, []int{1}, 0},
		{1, []int{1, 3, 5}, 0},
		{3, []int{1, 3, 5}, 1},
		{5, []int{1, 3, 5}, 2},
		{0, []int{1, 3, 5}, -1},
		{2, []int{1, 3, 5}, -1},
		{4, []int{1, 3, 5}, -1},
		{6, []int{1, 3, 5}, -1},
		{1, []int{1, 3, 5, 7}, 0},
		{3, []int{1, 3, 5, 7}, 1},
		{5, []int{1, 3, 5, 7}, 2},
		{7, []int{1, 3, 5, 7}, 3},
		{0, []int{1, 3, 5, 7}, -1},
		{2, []int{1, 3, 5, 7}, -1},
		{4, []int{1, 3, 5, 7}, -1},
		{6, []int{1, 3, 5, 7}, -1},
		{8, []int{1, 3, 5, 7}, -1},
	}
	for keyTest, test := range tests {
		t.Run("Testing iterative chop...", func(t *testing.T) {
			// run the test
			resChop := chop(test.elem, test.arr)
			if test.expectedIndex != chop(test.elem, test.arr) {
				t.Errorf("Test #%d failed; Received %d instead of %d", keyTest, resChop, test.expectedIndex)
			}
		})
	}

}
