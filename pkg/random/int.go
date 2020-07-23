package random

import (
	"math/rand"
	"time"
)

func (m *mizer) InitInt() {
	m.param = randInt(m.r)
}

func randInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
