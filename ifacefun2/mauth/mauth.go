package mauth

// Mauth tracks valid string values as a set.
type Mauth struct {
	m map[string]struct{}
}

// New receives strings and adds them to an instance of *Mauth as a set.
func New(ss ...string) *Mauth {
	m := make(map[string]struct{})

	for _, s := range ss {
		m[s] = struct{}{}
	}

	return &Mauth{
		m: m,
	}
}

// IsAuthorized verifies if a string exists in the tracked set.
func (m *Mauth) IsAuthorized(s string) bool {
	_, ok := m.m[s]
	return ok
}
