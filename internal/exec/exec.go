package exec

import (
	"io"

	"github.com/nurtai325/qurtc/internal/machine"
	"github.com/nurtai325/qurtc/internal/parser"
)

func Exec(stdout io.Writer, filename string, source []byte) error {
	newParser := parser.New(filename, source)
	decls, err := newParser.Parse()
	if err != nil {
		return err
	}
	program, err := machine.New(stdout, decls)
	if err != nil {
		return err
	}
	return program.Run()
}
