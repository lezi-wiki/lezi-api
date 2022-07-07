package services

import (
	"math/rand"
	"time"
)

func RandomInit() {
	rand.Seed(time.Now().UnixNano())
}

func Random(max int) int {
	RandomInit()
	randre := rand.Intn(max - 1)
	return randre + 1
}
