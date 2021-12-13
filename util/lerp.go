// Package util provides ...
package util

func Lerp(value int, min int, max int) int {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}
