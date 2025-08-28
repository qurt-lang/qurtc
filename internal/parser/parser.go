package parser

import (
	"errors"

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
Loop:
	for p.s.Scan() {
		switch p.s.Tok() {
		case token.FUNC:
			decls, err = p.appendDecl(decls, p.funcDecl)
			if err != nil {
				err = p.errorAt(errors.Join(err, ErrInvalidFuncDecl), help.FunctionsPage)
				break Loop
			}
		case token.STRUCT:
			decls, err = p.appendDecl(decls, p.structDecl)
			if err != nil {
				err = p.errorAt(errors.Join(err, ErrInvalidStructDecl), help.StructsPage)
				break Loop
			}
		case token.VAR:
			decls, err = p.appendDecl(decls, p.varDecl)
			if err != nil {
				err = p.errorAt(errors.Join(err, ErrInvalidVarDecl), help.VarsPage)
				break Loop
			}
		default:
			return nil, p.errorAt(ErrUnknownDecl, help.SyntaxPage)
		}
	}
	if p.s.Tok() == token.ILLEGAL {
		return nil, p.errorAt(p.s.Err(), help.QurtTour)
	} else if err != nil {
		return nil, err
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

func (p *parser) varDecl() (ast.Decl, error) {
	name, err := p.name()
	if err != nil {
		return nil, err
	}
	typ, err := p.typ()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.ASSIGN)
	if err != nil {
		return nil, err
	}
	val, err := p.expr()
	if err != nil {
		return nil, err
	}
	return &ast.VarDecl{
		Name: name,
		Type: typ,
		Val:  val,
	}, nil
}

func (p *parser) expr() (ast.Expr, error)

func (p *parser) typ() (ast.Type, error)

func (p *parser) name() (*ast.NameExpr, error) {
	lit, err := p.expect(token.IDENT)
	if err != nil {
		return nil, err
	}
	return &ast.NameExpr{Value: lit}, nil
}

func (p *parser) expect(tok token.Token) (string, error) {
	if !p.s.Scan() {
		return "", p.s.Err()
	} else if tok != p.s.Tok() {
		return "", ErrUnexpectedToken
	}
	return p.s.Lit(), nil
}
