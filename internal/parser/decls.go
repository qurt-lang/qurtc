package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) funcDecl() (ast.Decl, error) {
	typ, err := p.typ()
	if err != nil {
		return nil, err
	}
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
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok == token.RPAREN {
			p.expect(token.RPAREN)
			break
		}
		name, typ, err := p.fieldOrArg(token.RPAREN)
		if err != nil {
			return nil, err
		}
		args = append(args, &ast.FuncArg{
			Name: name,
			Type: typ,
		})
	}
	body, err := p.block()
	if err != nil {
		return nil, err
	}
	return &ast.FuncDecl{
		Name:       name,
		Args:       args,
		ReturnType: typ,
		Body:       body,
	}, nil
}

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
		name, typ, err := p.fieldOrArg(token.RBRACE)
		if err != nil {
			return nil, err
		}
		fields = append(fields, &ast.Field{
			Name: name,
			Type: typ,
		})
	}
	return &ast.StructDecl{
		Name:   name,
		Fields: fields,
	}, nil
}

func (p *parser) fieldOrArg(end token.Token) (string, *ast.Type, error) {
	name, err := p.name()
	if err != nil {
		return "", nil, err
	}
	typ, err := p.typ()
	if err != nil {
		return "", nil, err
	}
	tok, err := p.peek()
	if err != nil {
		return "", nil, err
	}
	if tok == token.COMMA {
		p.expect(token.COMMA)
		return name.Value, typ, nil
	} else if tok == end {
		return name.Value, typ, nil
	}
	return "", nil, ErrInvalidFieldOrArg
}
