package machine

import (
	"fmt"

	"github.com/nurtai325/qurtc/internal/ast"
)

func builtinFuncs(m *machine) map[string]*ast.BuiltinFuncDecl {
	builtinPrintName, builtinPrint := builtinPrint(m)
	return map[string]*ast.BuiltinFuncDecl{
		builtinPrintName: builtinPrint,
	}
}

func builtinPrint(m *machine) (string, *ast.BuiltinFuncDecl) {
	name := "жаз"
	return name, &ast.BuiltinFuncDecl{
		Name: &ast.NameExpr{Value: name},
		Body: func(args ...any) error {
			_, err := fmt.Fprintln(m.stdout, args)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
