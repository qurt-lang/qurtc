package machine

import (
	"fmt"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/types"
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
			_, err := fmt.Fprintln(m.stdout, args...)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func typesToAny(a []types.Type) []any {
	anys := make([]any, 0, len(a))
	for _, notAny := range a {
		anys = append(anys, notAny)
	}
	return anys
}
