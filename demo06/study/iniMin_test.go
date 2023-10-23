package study

import "testing"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMin(t *testing.T) {
	result := IntMin(2, -2)
	if result != -2 {
		t.Errorf("IntMin(2,-2) = %d;want -2", result)
	}
}
