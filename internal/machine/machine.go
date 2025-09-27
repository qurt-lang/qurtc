package machine

import (
	"fmt"
	"io"

	"github.com/nurtai325/qurtc/internal/ast"
)

const mainName = "негізгі"

type machine struct {
	stdout       io.Writer
	structs      map[string]*ast.StructDecl
	funcs        map[string]*ast.FuncDecl
	builtinFuncs map[string]*ast.BuiltinFuncDecl
}

func New(stdout io.Writer, decls []ast.Decl) (*machine, error) {
	mch := machine{
		stdout:  stdout,
		structs: make(map[string]*ast.StructDecl),
		funcs:   make(map[string]*ast.FuncDecl),
	}
	mch.builtinFuncs = builtinFuncs(&mch)
	for _, decl := range decls {
		switch v := decl.(type) {
		case *ast.StructDecl:
			if _, ok := mch.structs[v.Name.Value]; ok {
				return nil, fmt.Errorf("%w: %s", ErrDuplicateStruct, v.Name.Value)
			}
			mch.structs[v.Name.Value] = v
		case *ast.FuncDecl:
			if _, ok := mch.funcs[v.Name.Value]; ok {
				return nil, fmt.Errorf("%w: %s", ErrDuplicateFunc, v.Name.Value)
			}
			mch.funcs[v.Name.Value] = v
		}
	}
	return &mch, nil
}

func (m *machine) Run() error {
	main, ok := m.funcs[mainName]
	if !ok {
		return ErrNoMain
	}
	if len(main.Args) != 0 || main.ReturnType.Kind != ast.TVoid {
		return ErrInvalidMain
	}
	_, err := m.call(main.Name, nil)
	if err != nil {
		return err
	}
	return nil
}

func (m *machine) isFunc(expr ast.Expr) (*ast.FuncDecl, error) {
	funcName := expr.(*ast.NameExpr)
	return m.funcs[funcName.Value], nil
}

func (m *machine) isBuiltinFunc(expr ast.Expr) (*ast.BuiltinFuncDecl, error) {
	funcName := expr.(*ast.NameExpr)
	return m.builtinFuncs[funcName.Value], nil
}
