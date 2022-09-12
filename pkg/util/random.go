package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func RandomInt(min, max int) int {
	if min >= max {
		return max
	}
	return min + rand.Intn(max-min)
}

func RandomItemFromSlice[T any](slice []T) T {
	return slice[RandomInt(0, len(slice)-1)]
}
