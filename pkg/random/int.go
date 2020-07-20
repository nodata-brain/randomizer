package random

import ()

func (m mizer) InitInt() error {
	m.err = m.random()
	return nil
}
