package random

import ()

type mizer struct {
	param interface{}
	i     int
}

func Randomizer(f string, r int) (interface{}, error) {
	m := mizer{}
	m.setRange(r)
	err := m.random(f)
	if err != nil {
		return nil, err
	}
	return m.param, nil
}

func (m *mizer) random(f string) error {
	var err error
	switch f {
	case "int":
		m.InitInt()
	case "rnum":
		err = m.InitRnum()
	case "str":
		err = m.InitStr()
	default:
		m.InitInt()
	}
	return err
}

func (m *mizer) setRange(r int) {
	if r == 0 {
		m.randomRange()
	} else {
		m.i = r
	}
}
func (m *mizer) randomRange() {
	m.i = 10
}
