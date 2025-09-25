package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) varDecl() (ast.Decl, error) {
	varStmt, err := p.varStmt()
	if err != nil {
		return nil, err
	}
	return &ast.VarDecl{
		VarStmt: varStmt,
	}, nil
}

func (p *parser) structDecl() (ast.Decl, error) {
	name, err := p.name()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.LBRACE)
	if err != nil {
		return nil, err
	}
	var fields []*ast.Field
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok == token.RBRACE {
			p.expect(token.RBRACE)
			break
		}
		field, err := p.field()
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}
	return &ast.StructDecl{
		Name:   name,
		Fields: fields,
	}, nil
}

func (p *parser) field() (*ast.Field, error) {
	name, err := p.name()
	if err != nil {
		return nil, err
	}
	typ, err := p.typ()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.SEMICOLON)
	if err != nil {
		return nil, err
	}
	return &ast.Field{
		Name: name.Value,
		Type: typ,
	}, nil
}
