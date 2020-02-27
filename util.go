package multiprotocol

func ForEach(m Multiprotocol, cb func(c Component) bool) {
	// Shortcut if we already have a component
	if c, ok := m.(*Component); ok {
		cb(*c)
		return
	}

	b := m.Bytes()
	for len(b) > 0 {
		n, c, err := readComponent(b)
		if err != nil {
			panic(err)
		}
		if !cb(c) {
			return
		}
		b = b[n:]
	}
}

