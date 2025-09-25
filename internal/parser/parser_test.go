package parser_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/parser"
)

func TestParser(t *testing.T) {
	const examplesDir = "../../examples"
	entries, err := os.ReadDir(examplesDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		contents, err := os.ReadFile(fmt.Sprintf("%s/%s", examplesDir, entry.Name()))
		if err != nil {
			t.Fatal(err)
		}
		newParser := parser.New(entry.Name(), contents)
		decls, err := newParser.Parse()
		if err != nil {
			t.Errorf("expected successfully parsed file %s, got err %v", entry.Name(), err)
		}
		for _, decl := range decls {
			varDecl, ok := decl.(*ast.VarDecl)
			if !ok {
				structDecl, ok := decl.(*ast.StructDecl)
				if !ok {
					funcDecl := decl.(*ast.FuncDecl)
					fmt.Printf("%+v\n", funcDecl.Name)
					fmt.Printf("%+v\n", funcDecl.ReturnType)
					for _, arg := range funcDecl.Args {
						fmt.Printf("%s %+v\n", arg.Name, arg.Type)
					}
					for _, stmt := range funcDecl.Body {
						fmt.Printf("%+v\n", stmt)
					}
					fmt.Println()
					continue
				}
				fmt.Printf("%+v\n", structDecl.Name)
				for _, field := range structDecl.Fields {
					fmt.Printf("%s %+v\n", field.Name, field.Type)
				}
				fmt.Println()
				continue
			}
			fmt.Println(varDecl.Name.Value)
			fmt.Println(varDecl.Type.Name.Value)
			fmt.Println(varDecl.Type.Kind.String())
			fmt.Println(varDecl.Type.IsArray)
			fmt.Println(varDecl.Type.ArrayLen)

			switch v := varDecl.Val.(type) {
			case *ast.ArrayExpr:
				for _, el := range v.Elements {
					fmt.Printf("%+v, ", el)
				}
			case *ast.CallExpr:
				fmt.Println(v.Func)
				for _, el := range v.Args {
					fmt.Printf("%+v, ", el)
				}
			case *ast.OpExpr:
				fmt.Printf("%v ", v.Op)
				fmt.Printf("%+v ", v.Left)
				fmt.Printf("%+v\n", v.Right)
			default:
				fmt.Printf("%+v\n", varDecl.Val)
			}

			fmt.Printf("%T\n", varDecl.Val)
			fmt.Println()
		}
		t.Fail()
		// for _, decl := range decls {
		// 	funcDecl := decl.(*ast.FuncDecl)
		// 	fmt.Println(funcDecl.Name)
		// 	for _, arg := range funcDecl.Args {
		// 		fmt.Println("arg", arg.Name, arg.Type)
		// 	}
		// 	fmt.Println(funcDecl.ReturnType.Name)
		// 	for _, stmt := range funcDecl.Body {
		// 		fmt.Println(stmt)
		// 	}
		// }
	}
}
