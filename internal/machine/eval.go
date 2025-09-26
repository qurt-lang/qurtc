package machine

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/types"
)

func (m *machine) eval(exprScope *scope, expr ast.Expr) (types.Type, error) {
	switch v := expr.(type) {
	case *ast.StringExpr:
		return types.String(v.Value), nil
	case *ast.IntExpr:
		return types.Int(v.Value), nil
	case *ast.FloatExpr:
		return types.Float(v.Value), nil
	case *ast.BoolExpr:
		return types.Bool(v.Value), nil
	}
	return nil, nil
}

func (m *machine) evalAll(exprScope *scope, exprs []ast.Expr) ([]types.Type, error) {
	var evalled []types.Type
	for _, expr := range exprs {
		val, err := m.eval(exprScope, expr)
		if err != nil {
			return nil, err
		}
		evalled = append(evalled, val)
	}
	return evalled, nil
}

func (m *machine) call(fn ast.Expr, args []types.Type) (*ast.Expr, error) {
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
			return nil, builtinFunc.Body(args...)
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
