package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) varStmt() (*ast.VarStmt, error) {
	varName, err := p.name()
	if err != nil {
		return nil, err
	}
	varType, err := p.typ()
	if err != nil {
		return nil, err
	}
	if _, err = p.expect(token.ASSIGN); err != nil {
		return nil, err
	}
	val, err := p.expr(0)
	if err != nil {
		return nil, err
	}
	if _, err = p.expect(token.SEMICOLON); err != nil {
		return nil, err
	}
	return &ast.VarStmt{
		Name: varName,
		Type: varType,
		Val:  val,
	}, nil
}
