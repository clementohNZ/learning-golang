package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data []int
		answer int
	}

	tests := []test {
		test{[]int{21,21}, 42},
		test{[]int{4,5}, 9},
		test{[]int{1,1,3}, 5},
		test{[]int{-1,0,1}, 0},
	}

	for _,v := range tests {
		sum := mySum(v.data...)
		if sum != v.answer {
			t.Errorf("Expected %d, got %d", v.answer, sum)
		}
	}
}
