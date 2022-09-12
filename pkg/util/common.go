package util

import "math/rand"

func RandomInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return min + rand.Intn(max-min)
}

func RandomItemFromSlice[T any](slice []T) T {
	return slice[RandomInt(0, len(slice)-1)]
}
