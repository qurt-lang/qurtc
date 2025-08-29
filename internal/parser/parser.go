package parser

import (
	"errors"
	"strconv"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/help"
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
	// TODO: add known errors, return them from functions without panicking
	for p.s.Scan() {
		switch p.s.Tok() {
		case token.FUNC:
			decls, err = p.appendDecl(decls, p.funcDecl)
			if err != nil {
				return nil, p.errorAt(errors.Join(err, ErrInvalidFuncDecl), help.FunctionsPage)
			}
		case token.STRUCT:
			decls, err = p.appendDecl(decls, p.structDecl)
			if err != nil {
				return nil, p.errorAt(errors.Join(err, ErrInvalidStructDecl), help.StructsPage)
			}
		case token.VAR:
			decls, err = p.appendDecl(decls, p.varDecl)
			if err != nil {
				return nil, p.errorAt(errors.Join(err, ErrInvalidVarDecl), help.VarsPage)
			}
		default:
			return nil, p.errorAt(ErrUnknownDecl, help.SyntaxPage)
		}
	}
	if p.s.Tok() == token.ILLEGAL {
		return nil, p.errorAt(p.s.Err(), help.QurtTour)
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

func (p *parser) name() (*ast.NameExpr, error) {
	lit, err := p.expect(token.IDENT)
	if err != nil {
		return nil, err
	}
	return &ast.NameExpr{Value: lit}, nil
}

func (p *parser) typ() (*ast.Type, error) {
	_, tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	var parsedType ast.Type
	if tok == token.LBRACK {
		_, err = p.expect(token.LBRACK)
		if err != nil {
			return nil, err
		}
		lit, err := p.expect(token.INT)
		if err != nil {
			return nil, err
		}
		_, err = p.expect(token.RBRACK)
		if err != nil {
			return nil, err
		}
		arrayLen, err := strconv.ParseInt(lit, 10, 0)
		if err != nil {
			return nil, ErrUnexpectedToken
		}
		parsedType.IsArray = true
		parsedType.ArrayLen = int(arrayLen)
	}
	name, err := p.name()
	if err != nil {
		return nil, err
	}
	parsedType.Name = name
	return &parsedType, nil
}

func (p *parser) expect(tok token.Token) (string, error) {
	if !p.s.Scan() {
		return "", p.s.Err()
	} else if tok != p.s.Tok() {
		return "", ErrUnexpectedToken
	}
	return p.s.Lit(), nil
}

func (p *parser) peek() (string, token.Token, error) {
	if !p.s.Scan() {
		return "", 0, p.s.Err()
	}
	p.s.Back()
	return p.s.Lit(), p.s.Tok(), nil
}
