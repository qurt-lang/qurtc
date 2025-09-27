package types

func NewStruct(name string, fields map[string]Type) (*Struct, error) {
	return &Struct{
		name:   name,
		fields: fields,
	}, nil
}

func (a *Struct) Get(name string) (Type, error) {
	field, ok := a.fields[name]
	if !ok {
		return nil, ErrNoSuchField
	}
	return field, nil
}

func (a *Struct) Set(name string, val Type) error {
	_, ok := a.fields[name]
	if !ok {
		return ErrNoSuchField
	} else if !IsSameType(a.fields[name], val) {
		return ErrNotSameType
	}
	a.fields[name] = val
	return nil
}
