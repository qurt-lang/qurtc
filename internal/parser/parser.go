package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/scanner"
	"github.com/nurtai325/qurtc/internal/token"
)

type parser struct {
	s scanner.Scanner
}

func New(filename string, input []byte) *parser {
	return &parser{
		s: scanner.New(filename, []byte(input)),
	}
}

func (p *parser) Parse() (decls []ast.Decl, err error) {
	for p.s.Scan() {
		switch p.s.Tok() {
		case token.STRUCT:
			decls, err = p.appendDecl(decls, p.structDecl)
			if err != nil {
				return nil, err
			}

		case token.FUNC:
			decls, err = p.appendDecl(decls, p.funcDecl)
			if err != nil {
				return nil, err
			}

		default:
			return nil, p.errorAt(ErrUnknownDecl, "")
		}
	}

	if p.s.Tok() == token.ILLEGAL {
		return nil, p.errorAt(p.s.Err(), "")
	}

	return decls, nil
}

func (p *parser) appendDecl(decls []ast.Decl, fn func() (ast.Decl, error)) ([]ast.Decl, error) {
	decl, err := fn()
	if err != nil {
		return nil, err
	}

	return append(decls, decl), nil
}

func (p *parser) structDecl() (ast.Decl, error) {
	return nil, nil
}

func (p *parser) funcDecl() (ast.Decl, error) {
	return nil, nil
}

func (p *parser) name() (*ast.NameExpr, error) {
	lit, err := p.expect(token.IDENT)
	if err != nil {
		return nil, err
	}

	return &ast.NameExpr{Value: lit}, nil
}

func (p *parser) typ() (*ast.Expr, error)

func (p *parser) expect(tok token.Token) (string, error) {
	if !p.s.Scan() {
		return "", p.s.Err()
	} else if tok != p.s.Tok() {
		return "", ErrUnexpectedToken
	}
	return p.s.Lit(), nil
}
