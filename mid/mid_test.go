package mid

import "testing"

func TestCalculate(t *testing.T) {
	t.Log(calculate("5+(2-1)*4"))
}

func TestGaryCode(t *testing.T) {
	t.Log(grayCode(3))
}

func TestNumDecodings(t *testing.T) {
	t.Log(numDecodings("226"))
}
func TestPermutation(t *testing.T) {
	for _, str := range permutation("eqq") {
		t.Log(str)
	}
}
