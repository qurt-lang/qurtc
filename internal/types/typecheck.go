package types

import (
	"fmt"

	"github.com/nurtai325/qurtc/internal/ast"
)

func ZeroOf(typ *ast.Type) (Type, error) {
	switch typ.Kind {
	case ast.TInt:
		return Int(0), nil
	case ast.TFloat:
		return Float(0), nil
	case ast.TString:
		return String(""), nil
	case ast.TBool:
		return Bool(false), nil
	default:
		return nil, ErrUnknownType
	}
}

func IsOfType(val Type, typ *ast.Type) bool {
	switch v := val.(type) {
	case Int:
		return typ.Kind == ast.TInt
	case Float:
		return typ.Kind == ast.TFloat
	case String:
		return typ.Kind == ast.TString
	case Bool:
		return typ.Kind == ast.TBool
	case *Array:
		if !typ.IsArray {
			return false
		} else if typ.ArrayLen != v.Len() {
			return false
		}
		for _, el := range v.elements {
			if !IsOfType(el, typ) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func IsSameType(vals ...Type) bool {
	var typeName string
	for _, val := range vals {
		if typeName == "" {
			typeName = fmt.Sprintf("%T", val)
		} else if typeName != fmt.Sprintf("%T", val) {
			return false
		}
	}
	return true
}
