package parser_test

import (
	"testing"

	"github.com/nurtai325/qurtc/internal/parser"
	"github.com/nurtai325/qurtc/internal/testutils"
)

func TestParser(t *testing.T) {
	testutils.RunOnExamples(func(name string, contents []byte) {
		newParser := parser.New(name, contents)
		_, err := newParser.Parse()
		if err != nil {
			t.Errorf("expected successfully parsed file %s, got err %v", name, err)
		}
	})
}
