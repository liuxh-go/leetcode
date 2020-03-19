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
