package machine

import (
	"fmt"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/parser"
	"github.com/nurtai325/qurtc/internal/token"
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
	case *ast.NameExpr:
		variable := exprScope.get(v.Value)
		if variable == nil {
			return nil, ErrUndefinedReference
		}
		return variable, nil
	case *ast.ArrayExpr:
		fmt.Println("hello")
		elements, err := m.evalAll(exprScope, v.Elements)
		if err != nil {
			fmt.Println("hello")
			return nil, err
		}
		if !types.IsSameType(elements...) {
			return nil, types.ErrNotSameType
		}
		return types.NewArray(elements)
	case *ast.CallExpr:
		args, err := m.evalAll(exprScope, v.Args)
		if err != nil {
			return nil, err
		}
		return m.call(v.Func, args)
	case *ast.OpExpr:
		left, err := m.eval(exprScope, v.Left)
		if err != nil {
			return nil, err
		}
		right, err := m.eval(exprScope, v.Right)
		if err != nil {
			return nil, err
		}
		res, err := m.binary(v.Op, left, right)
		if err != nil {
			return nil, err
		}
		return res, nil
	case *ast.UnaryOpExpr:
		operand, err := m.eval(exprScope, v.Operand)
		if err != nil {
			return nil, err
		}
		return m.unary(v.Op, operand)
	default:
		return nil, parser.ErrInvalidExpr
	}
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

func (m *machine) call(fn ast.Expr, args []types.Type) (types.Type, error) {
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
			return nil, builtinFunc.Body(typesToAny(args)...)
		}
		return nil, ErrCallNoFunc
	}

	currScope, err := newFuncScope(funcDecl, args)
	if err != nil {
		return nil, err
	}
	for _, stmt := range funcDecl.Body {
		retVal, err := m.exec(currScope, stmt)
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

func (m *machine) binary(op token.Token, x, y types.Type) (types.Type, error) {
	if !types.IsSameType(x, y) {
		return nil, ErrNotSameTypeOp
	}
	switch op {
	case token.MUL:
		return m.mul(x, y)
	case token.DIV:
		return m.div(x, y)
	case token.MOD:
		return m.mod(x, y)
	case token.ADD:
		return m.add(x, y)
	case token.SUB:
		return m.sub(x, y)
	case token.EQL:
		return m.eql(x, y)
	case token.LSS:
		return m.lss(x, y)
	case token.GTR:
		return m.gtr(x, y)
	case token.NEQ:
		return m.neq(x, y)
	case token.LEQ:
		return m.leq(x, y)
	case token.GEQ:
		return m.geq(x, y)
	case token.LAND:
		return m.land(x, y)
	case token.LOR:
		return m.lor(x, y)
	default:
		return nil, ErrUnknownOp
	}
}

func (m *machine) unary(op token.Token, x types.Type) (types.Type, error) {
	switch op {
	case token.NOT:
		switch x := x.(type) {
		case types.Bool:
			return !x, nil
		default:
			return nil, ErrOpNotSupportedForType
		}
	case token.SUB:
		switch x := x.(type) {
		case types.Int:
			return -x, nil
		case types.Float:
			return -x, nil
		default:
			return nil, ErrOpNotSupportedForType
		}
	default:
		return nil, ErrUnknownOp
	}
}

func (m *machine) mul(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return x * y, nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return x * y, nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) div(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return x / y, nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return x / y, nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) mod(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return x % y, nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) add(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return x + y, nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return x + y, nil
	case types.String:
		x, y := x.(types.String), y.(types.String)
		return x + y, nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) sub(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return x - y, nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return x - y, nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) eql(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return types.Bool(x == y), nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return types.Bool(x == y), nil
	case types.String:
		x, y := x.(types.String), y.(types.String)
		return types.Bool(x == y), nil
	case types.Bool:
		x, y := x.(types.Bool), y.(types.Bool)
		return types.Bool(x == y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) lss(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return types.Bool(x < y), nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return types.Bool(x < y), nil
	case types.String:
		x, y := x.(types.String), y.(types.String)
		return types.Bool(x < y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) gtr(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return types.Bool(x > y), nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return types.Bool(x > y), nil
	case types.String:
		x, y := x.(types.String), y.(types.String)
		return types.Bool(x > y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) neq(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return types.Bool(x != y), nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return types.Bool(x != y), nil
	case types.String:
		x, y := x.(types.String), y.(types.String)
		return types.Bool(x != y), nil
	case types.Bool:
		x, y := x.(types.Bool), y.(types.Bool)
		return types.Bool(x != y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) leq(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return types.Bool(x <= y), nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return types.Bool(x <= y), nil
	case types.String:
		x, y := x.(types.String), y.(types.String)
		return types.Bool(x <= y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) geq(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Int:
		x, y := x.(types.Int), y.(types.Int)
		return types.Bool(x >= y), nil
	case types.Float:
		x, y := x.(types.Float), y.(types.Float)
		return types.Bool(x >= y), nil
	case types.String:
		x, y := x.(types.String), y.(types.String)
		return types.Bool(x >= y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) land(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Bool:
		x, y := x.(types.Bool), y.(types.Bool)
		return types.Bool(x && y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}

func (m *machine) lor(x, y types.Type) (types.Type, error) {
	switch x.(type) {
	case types.Bool:
		x, y := x.(types.Bool), y.(types.Bool)
		return types.Bool(x || y), nil
	default:
		return nil, ErrOpNotSupportedForType
	}
}
