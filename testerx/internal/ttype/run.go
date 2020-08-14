package ttype

type Skipper interface {
	Skip()
}

func RunForTypes(s Skipper, ts ...TType) {
	if Running == All {
		return
	}

	for _, t := range ts {
		if t == Running {
			return
		}
	}

	s.Skip()
}
