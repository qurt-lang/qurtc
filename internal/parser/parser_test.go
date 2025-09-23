package parser_test

import (
	"fmt"
	"os"
	"testing"

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
		_, err = newParser.Parse()
		if err != nil {
			t.Errorf("expected successfully parsed file %s, got err %v", entry.Name(), err)
		}
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
