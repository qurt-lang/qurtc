package machine

import "github.com/nurtai325/qurtc/internal/ast"

func (m *machine) isFunc(expr ast.Expr) (*ast.FuncDecl, error) {
	funcName := expr.(*ast.NameExpr)
	return m.funcs[funcName.Value], nil
}

func (m *machine) isBuiltinFunc(expr ast.Expr) (*ast.BuiltinFuncDecl, error) {
	funcName := expr.(*ast.NameExpr)
	return m.builtinFuncs[funcName.Value], nil
}
