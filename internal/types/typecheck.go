package types

import "github.com/nurtai325/qurtc/internal/ast"

func ZeroOf(typ *ast.Type) Type {
	switch typ.Kind {
	case ast.TInt:
		return Int(0)
	case ast.TFloat:
		return Float(0)
	case ast.TString:
		return String("")
	case ast.TBool:
		return Bool(false)
	default:
		return nil
	}
}

func IsOfType(val Type, arg *ast.FuncArg) bool {
	switch arg.Type.Kind {
	case ast.TInt:
		_, ok := val.(Int)
		return ok
	case ast.TFloat:
		_, ok := val.(Float)
		return ok
	case ast.TString:
		_, ok := val.(String)
		return ok
	case ast.TBool:
		_, ok := val.(Bool)
		return ok
	default:
		return false
	}
}
