package machine

type scope struct {
}

func (s *scope) with(childScope *scope) *scope {
	return nil
}
