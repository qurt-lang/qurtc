package machine

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/parser"
)

func (m *machine) exec(parentScope *scope, stmt ast.Stmt) (*ast.Expr, error) {
	switch v := stmt.(type) {
	case *ast.CallStmt:
		args, err := m.evalAll(parentScope, v.CallExpr.Args)
		if err != nil {
			return nil, err
		}
		return m.call(v.CallExpr.Func, args)
	default:
		return nil, parser.ErrUnknownStmt
	}
}
