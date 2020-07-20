package random

import ()

func (m *mizer) InitStr() error {
	m.err = m.random()
	return nil
}
