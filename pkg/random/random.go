package random

import ()

type mizer struct {
	param interface{}
	max   int
	min   int
	err   error
}

func Randomizer(f string, max int, min int) interface{} {
	m := mizer{}
	m.setRange(max, min)
	switch f {
	case "int":
		m.err = m.InitInt()
	case "string":
		m.err = m.InitStr()
	default:
		m.err = nil
	}
	return m.param
}

func (m mizer) random() error {
	return nil
}

func (m *mizer) setRange(max interface{}, min interface{}) {
	if min == nil {
		m.randomRange()
	} else {
	}

}
func (m *mizer) randomRange() {
	m.min = 1
	m.max = 5
}
