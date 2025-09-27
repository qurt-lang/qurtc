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
			res, err := types.ZeroOf(v.Type, m.structs)
			if err != nil {
				return nil, err
			}
			val = res
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
	case *ast.AssignStmt:
		val, err := m.eval(parentScope, v.Val)
		if err != nil {
			return nil, err
		}
		switch assignee := v.Var.(type) {
		case *ast.NameExpr:
			if !types.IsSameType(parentScope.get(assignee.Value), val) {
				return nil, types.ErrNotSameType
			}
			if !parentScope.set(assignee.Value, val) {
				return nil, types.ErrNotSameType
			}
			return nil, nil
		case *ast.ArrayAccessExpr:
			res, err := m.eval(parentScope, assignee.Array)
			if err != nil {
				return nil, err
			}
			arr, ok := res.(*types.Array)
			if !ok {
				return nil, ErrArrAccessOnNotArr
			}
			res, err = m.eval(parentScope, assignee.Index)
			if err != nil {
				return nil, err
			}
			index, ok := res.(types.Int)
			if !ok {
				return nil, ErrArrAccessOnNotArr
			}
			return nil, arr.Set(int(index), val)
		case *ast.SelectorExpr:
			res, err := m.eval(parentScope, assignee.Struct)
			if err != nil {
				return nil, err
			}
			structVal, ok := res.(*types.Struct)
			if !ok {
				return nil, ErrArrAccessOnNotArr
			}
			return nil, structVal.Set(assignee.Field.Value, val)
		default:
			return nil, ErrInvalidAssign
		}
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
