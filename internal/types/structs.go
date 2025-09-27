package types

func NewStruct(name string, fields map[string]Type) (*Struct, error) {
	return &Struct{
		typeName: name,
		fields:   fields,
	}, nil
}

func (s *Struct) Get(name string) (Type, error) {
	field, ok := s.fields[name]
	if !ok {
		return nil, ErrNoSuchField
	}
	return field, nil
}

func (s *Struct) Set(name string, val Type) error {
	_, ok := s.fields[name]
	if !ok {
		return ErrNoSuchField
	} else if !IsSameType(s.fields[name], val) {
		return ErrNotSameType
	}
	s.fields[name] = val
	return nil
}
