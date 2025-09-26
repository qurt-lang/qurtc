package types

import "github.com/nurtai325/qurtc/internal/token"

type (
	Type interface {
		aType()
	}

	Int int

	Float float32

	String string

	Bool bool
)

func (b Bool) String() string {
	if b {
		return token.TRUE.String()
	} else {
		return token.FALSE.String()
	}
}

func (Int) aType() {}

func (Float) aType() {}

func (String) aType() {}

func (Bool) aType() {}
