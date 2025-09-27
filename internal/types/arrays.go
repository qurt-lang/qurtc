package types

func NewArray(elements []Type) (*Array, error) {
	return &Array{
		elements: elements,
		length:   len(elements),
	}, nil
}

func (a *Array) Len() int {
	return a.length
}

func (a *Array) Get(i int) (Type, error) {
	if a.isOutOfBound(i) {
		return nil, ErrOutOfBound
	}
	return a.elements[i], nil
}

func (a *Array) Set(i int, val Type) error {
	if a.isOutOfBound(i) {
		return ErrOutOfBound
	} else if !IsSameType(a.elements[i], val) {
		return ErrNotSameType
	}
	a.elements[i] = val
	return nil
}

func (a *Array) isOutOfBound(i int) bool {
	if a.length <= i || i < 0 {
		return true
	}
	return false
}
