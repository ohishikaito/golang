package mylib

import "testing"

func TestAverage(t *testing.T) {
	t.Skip("Hoge")
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 4 {
		t.Error(v)
	}
}
