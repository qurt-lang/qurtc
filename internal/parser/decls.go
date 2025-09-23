package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) varDecl() (ast.Decl, error) {
	varName, err := p.name()
	if err != nil {
		return nil, err
	}
	varType, err := p.typ()
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
	_, err = p.expect(token.SEMICOLON)
	if err != nil {
		return nil, err
	}
	return &ast.VarDecl{
		Name: varName,
		Type: varType,
		Val:  val,
	}, nil
}
