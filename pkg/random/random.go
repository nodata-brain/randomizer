package random

import ()

type mizer struct {
	param interface{}
	r     int
	err   error
}

func Randomizer(f string, r int) interface{} {
	m := mizer{}
	m.setRange(r)
	m.random(f)
	return m.param
}

func (m *mizer) random(f string) error {
	switch f {
	case "int":
		m.InitInt()
	case "str":
		m.err = m.InitStr()
	default:
		m.InitInt()
	}
	return nil
}

func (m *mizer) setRange(r int) {
	if r == 0 {
		m.randomRange()
	} else {
		m.r = r
	}
}
func (m *mizer) randomRange() {
	m.r = 10
}
