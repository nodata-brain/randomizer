package random

import (
	"math/rand"
	"strconv"
	"time"
)

func (m *mizer) InitInt() {
	m.param = randInt(m.i)
}

func randInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func (m *mizer) InitRnum() error {
	var err error
	s := []int{}
	for i := 0; m.i > i; i++ {
		if i == 0 {
			for {
				ra := randInt(9)
				if ra != 0 {
					s = append(s, ra)
					break
				}
			}
			continue
		}
		s = append(s, randInt(9))
	}
	var num string
	for _, n := range s {
		str := strconv.Itoa(n)
		num = num + str
	}
	m.param, err = strconv.Atoi(num)
	if err != nil {
		return err
	}
	return nil
}
