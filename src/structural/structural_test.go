package structural

import "testing"

func TestChop(t *testing.T) {
	var tests = []struct {
		elem          int
		cInput        *chopInput
		expectedIndex int
	}{
		{3, &chopInput{arr: []int{}}, -1},
		{3, &chopInput{arr: []int{1}}, -1},
		{1, &chopInput{arr: []int{1}}, 0},
		{1, &chopInput{arr: []int{1, 3, 5}}, 0},
		{3, &chopInput{arr: []int{1, 3, 5}}, 1},
		{5, &chopInput{arr: []int{1, 3, 5}}, 2},
		{0, &chopInput{arr: []int{1, 3, 5}}, -1},
		{2, &chopInput{arr: []int{1, 3, 5}}, -1},
		{4, &chopInput{arr: []int{1, 3, 5}}, -1},
		{6, &chopInput{arr: []int{1, 3, 5}}, -1},
		{1, &chopInput{arr: []int{1, 3, 5, 7}}, 0},
		{3, &chopInput{arr: []int{1, 3, 5, 7}}, 1},
		{5, &chopInput{arr: []int{1, 3, 5, 7}}, 2},
		{7, &chopInput{arr: []int{1, 3, 5, 7}}, 3},
		{0, &chopInput{arr: []int{1, 3, 5, 7}}, -1},
		{2, &chopInput{arr: []int{1, 3, 5, 7}}, -1},
		{4, &chopInput{arr: []int{1, 3, 5, 7}}, -1},
		{6, &chopInput{arr: []int{1, 3, 5, 7}}, -1},
		{8, &chopInput{arr: []int{1, 3, 5, 7}}, -1},
	}
	for keyTest, test := range tests {
		t.Run("Testing recursive chop...", func(t *testing.T) {
			// run the test
			resChop := test.cInput.chop(test.elem)
			if test.expectedIndex != resChop {
				t.Errorf("Test #%d failed; Received %d instead of %d", keyTest, resChop, test.expectedIndex)
			}
		})
	}

}
