package types

import (
	"fmt"

	"github.com/nurtai325/qurtc/internal/token"
)

type (
	Type interface {
		aType()
	}

	Int int

	Float float32

	String string

	Bool bool

	Array struct {
		elements []Type
		length   int
	}

	Struct struct {
		name   string
		fields map[string]Type
	}
)

func (b Bool) String() string {
	if b {
		return token.TRUE.String()
	} else {
		return token.FALSE.String()
	}
}

func (a *Array) String() string {
	return fmt.Sprint(a.elements)
}

func (Int) aType() {}

func (Float) aType() {}

func (String) aType() {}

func (Bool) aType() {}

func (*Array) aType() {}

func (*Struct) aType() {}
