package machine

import (
	"github.com/nurtai325/qurtc/internal/ast"
)

func (m *machine) eval(expr ast.Expr) (any, error) {
	return nil, nil
}

func (m *machine) call(fn ast.Expr, args ...ast.Expr) (*ast.Expr, error) {
	funcDecl, err := m.isFunc(fn)
	if err != nil {
		return nil, err
	}
	if funcDecl == nil {
		builtinFunc, err := m.isBuiltinFunc(fn)
		if err != nil {
			return nil, err
		}
		if builtinFunc != nil {
			return nil, builtinFunc.Body(args)
		}
		return nil, ErrCallNoFunc
	}

	var currScope scope
	for _, stmt := range funcDecl.Body {
		retVal, err := m.exec(&currScope, stmt)
		if err != nil {
			return nil, err
		}
		if retVal != nil {
			return retVal, nil
		}
	}
	// return nil because func body didn't return anything
	return nil, nil
}
