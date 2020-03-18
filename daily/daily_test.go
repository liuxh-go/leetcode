package daily

import "testing"

func TestIsRectangleOverlap(t *testing.T) {
	t.Log(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}
