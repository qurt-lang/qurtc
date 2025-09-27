package machine

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/parser"
	"github.com/nurtai325/qurtc/internal/types"
)

func (m *machine) exec(parentScope *scope, stmt ast.Stmt) (types.Type, error) {
	switch v := stmt.(type) {
	case *ast.VarStmt:
		var val types.Type
		if v.Val == nil {
			if v.Type.IsArray {
				elements := make([]types.Type, 0, v.Type.ArrayLen)
				for range v.Type.ArrayLen {
					val, err := types.ZeroOf(v.Type)
					if err != nil {
						return nil, err
					}
					elements = append(elements, val)
				}
				arr, err := types.NewArray(elements)
				if err != nil {
					return nil, err
				}
				val = arr
			} else {
				res, err := types.ZeroOf(v.Type)
				if err != nil {
					return nil, err
				}
				val = res
			}
		} else {
			res, err := m.eval(parentScope, v.Val)
			if err != nil {
				return nil, err
			}
			if !types.IsOfType(res, v.Type) {
				return nil, types.ErrNotSameType
			}
			val = res
		}
		if !parentScope.add(v.Name.Value, val) {
			return nil, ErrVarExists
		}
		return nil, nil
	case *ast.ReturnStmt:
		res, err := m.eval(parentScope, v.Value)
		if err != nil {
			return nil, err
		}
		return res, nil
	case *ast.CallStmt:
		args, err := m.evalAll(parentScope, v.CallExpr.Args)
		if err != nil {
			return nil, err
		}
		_, err = m.call(v.CallExpr.Func, args)
		if err != nil {
			return nil, err
		}
		return nil, nil
	default:
		return nil, parser.ErrUnknownStmt
	}
}
