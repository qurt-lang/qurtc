package machine_test

import (
	"os"
	"testing"

	"github.com/nurtai325/qurtc/internal/machine"
	"github.com/nurtai325/qurtc/internal/parser"
	"github.com/nurtai325/qurtc/internal/testutils"
)

func TestMachine(t *testing.T) {
	testutils.RunOnExamples(func(name string, contents []byte) {
		newParser := parser.New(name, contents)
		decls, err := newParser.Parse()
		if err != nil {
			t.Fatal(err)
		}
		program, err := machine.New(os.Stdout, decls)
		if err != nil {
			t.Fatal(err)
		}
		err = program.Run()
		if err != nil {
			t.Fatal(err)
		}
	})
}
