// Package custommath implements a simple math function
package custommath

// Sum adds an unlimited amount of values of type int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
