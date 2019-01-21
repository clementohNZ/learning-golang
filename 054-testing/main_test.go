package main

import "testing"

func TestMySum(t *testing.T) {
	sum := mySum(2,3)
	if sum != 5 {
		t.Errorf("Expected 5, got %d", sum)
	}
}
