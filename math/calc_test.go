package math

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(2, 5)
	if ans != 7 {
		t.Errorf("Add(2, 5) = %d; want 7", ans)
	}
}

func TestSubtractNums(t *testing.T) {
	ans := SubtractNums(10, 5)
	if ans != 5 {
		t.Errorf("SubtractNums(10, 5) = %d; want 5", ans)
	}
}
