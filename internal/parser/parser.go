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

func (p *parser) structDecl() (ast.Decl, error) {
	name, err := p.name()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.LBRACK)
	if err != nil {
		return nil, err
	}

	var fields []*ast.Field
	for {
		_, tok, err := p.peekNext()
		if err != nil {
			return nil, err
		} else if tok == token.RBRACE {
			break
		}

		fieldName, err := p.name()
		if err != nil {
			return nil, err
		}
		fieldType, err := p.typ()
		if err != nil {
			return nil, err
		}
		_, err = p.expect(token.SEMICOLON)
		if err != nil {
			return nil, err
		}
		fields = append(fields, &ast.Field{
			Name: fieldName.Value,
			Type: fieldType,
		})
	}
	_, err = p.expect(token.RBRACK)
	if err != nil {
		return nil, err
	}

	return &ast.StructDecl{
		Name:   name,
		Fields: fields,
	}, nil
}

func (p *parser) funcDecl() (ast.Decl, error) {
	name, err := p.name()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.LPAREN)
	if err != nil {
		return nil, err
	}

	var args []*ast.FuncArg
	for {
		_, tok, err := p.peekNext()
		if err != nil {
			return nil, err
		} else if tok == token.RPAREN {
			break
		}

		argName, err := p.name()
		if err != nil {
			return nil, err
		}
		argType, err := p.typ()
		if err != nil {
			return nil, err
		}
		_, err = p.expect(token.COMMA)
		if err != nil {
			return nil, err
		}
		args = append(args, &ast.FuncArg{
			Name: argName.Value,
			Type: argType,
		})
	}
	_, err = p.expect(token.RPAREN)
	if err != nil {
		return nil, err
	}

	returnType, err := p.typ()
	if err != nil {
		return nil, err
	}

	_, err = p.expect(token.LBRACE)
	if err != nil {
		return nil, err
	}
	var body []ast.Stmt
	for {
		_, tok, err := p.peekNext()
		if err != nil {
			return nil, err
		} else if tok == token.RBRACE {
			break
		}
		stmt, err := p.stmt()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}
	_, err = p.expect(token.RBRACE)
	if err != nil {
		return nil, err
	}
	return &ast.FuncDecl{
		Name:       name,
		Args:       args,
		ReturnType: returnType,
		Body:       body,
	}, nil
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

func (p *parser) typ() (*ast.Type, error)

func (p *parser) stmt() (ast.Stmt, error)

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

func (p *parser) peekNext() (string, token.Token, error) {
	if !p.s.Scan() {
		return "", 0, p.s.Err()
	}
	p.s.Back()
	return p.s.Lit(), p.s.Tok(), nil
}
