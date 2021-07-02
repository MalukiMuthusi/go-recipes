package main

// filter returns a slice with only the values that pred(val) returned true
func filter(pred func(int) bool, values []int) []int {
	vals := make([]int, 0)
	for _, v := range values {
		if pred(v) {
			vals = append(vals, v)
		}
	}
	return vals
}

func isOdd(n int) bool {
	return n%2 == 1
}
