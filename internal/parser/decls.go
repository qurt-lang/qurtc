package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

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
		_, tok, err := p.peek()
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
		_, tok, err := p.peek()
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
		_, tok, err := p.peek()
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
		_, err = p.expect(token.COMMA)
		if err != nil {
			return nil, err
		}
		fields = append(fields, &ast.Field{
			Name: fieldName.Value,
			Type: fieldType,
		})
	}
	_, err = p.expect(token.RBRACE)
	if err != nil {
		return nil, err
	}

	return &ast.StructDecl{
		Name:   name,
		Fields: fields,
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
