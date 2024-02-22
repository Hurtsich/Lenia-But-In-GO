package collection

import "fmt"

func Contains[S ~[]E, E any](s S, v E, comparer func(r, t E) bool) bool {
	return Index(s, v, comparer) >= 0
}

func Index[S ~[]E, E any](s S, v E, comparer func(r, t E) bool) int {
	for i := range s {
		if comparer(v, s[i]) {
			fmt.Printf("Comparing %v, %v\n", v, s[i])
			return i
		}
	}
	return -1
}
