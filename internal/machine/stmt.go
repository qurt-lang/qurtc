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
	case *ast.IfStmt:
		res, err := m.eval(parentScope, v.Cond)
		if err != nil {
			return nil, err
		}
		cond, ok := res.(types.Bool)
		if !ok {
			return nil, ErrIfWithNoBool
		}
		ifScope := parentScope.clone()
		if cond == false {
			if v.Else == nil {
				return nil, nil
			}
			switch elseBlock := v.Else.(type) {
			case ast.Stmts:
				return m.execBlock(ifScope.clone(), v.Then)
			case *ast.IfStmt:
				return m.exec(ifScope, elseBlock)
			default:
				return nil, ErrInvalidElse
			}
		}
		return m.execBlock(ifScope, v.Then)
	case *ast.ForStmt:
		loopScope := parentScope.clone()
		loopScope.isLoop = true
		_, err := m.exec(loopScope, v.Init)
		if err != nil {
			return nil, err
		}

		for {
			res, err := m.eval(loopScope, v.Cond)
			if err != nil {
				return nil, err
			}
			cond, ok := res.(types.Bool)
			if !ok {
				return nil, ErrIfWithNoBool
			}
			if !cond {
				break
			}

			retVal, err := m.execBlock(loopScope, v.Body)
			if err != nil {
				return nil, err
			}
			if retVal != nil {
				return retVal, nil
			}
			if loopScope.isBreak {
				break
			}

			_, err = m.exec(loopScope, v.Post)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case *ast.ContinueStmt:
		if !parentScope.isLoop {
			return nil, ErrContinueInNotLoop
		}
		parentScope.isContinue = true
		return nil, nil
	case *ast.BreakStmt:
		if !parentScope.isLoop {
			return nil, ErrBreakInNotLoop
		}
		parentScope.isBreak = true
		return nil, nil
	default:
		return nil, parser.ErrUnknownStmt
	}
}

func (m *machine) execBlock(currScope *scope, block []ast.Stmt) (types.Type, error) {
	for _, stmt := range block {
		retVal, err := m.exec(currScope, stmt)
		if err != nil {
			return nil, err
		}
		if retVal != nil {
			return retVal, nil
		}
		if currScope.isContinue {
			break
		}
	}
	// return nil because func body didn't return anything
	return nil, nil
}
