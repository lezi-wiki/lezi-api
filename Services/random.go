package Services

import (
	"math/rand"
	"time"
)

func RandomInit() {
	rand.Seed(time.Now().UnixNano())
}

func Random() int {
	RandomInit()
	randre := rand.Intn(100) //TODO Read database id then set random max
	return randre
}
