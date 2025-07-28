package parser

import (
	"github.com/nurtai325/qurt/internal/ast"
	"github.com/nurtai325/qurt/internal/scanner"
	"github.com/nurtai325/qurt/internal/token"
)

type parser struct {
	s scanner.Scanner
}

func New(filename string, input []byte) *parser {
	return &parser{
		s: scanner.New(filename, []byte(input)),
	}
}

func (p *parser) Parse() ([]ast.Decl, error) {
	var decls []ast.Decl

	for p.s.Scan() {
		lit, tok := p.s.Lit(), p.s.Tok()
		if tok == token.EOF {
			break
		} else if tok == token.ILLEGAL {
			return nil, p.lexError()
		}

		switch tok {

		case token.STRUCT:

		case token.VAR:

		case token.FUNC:

		}
	}

	return decls, nil
}
