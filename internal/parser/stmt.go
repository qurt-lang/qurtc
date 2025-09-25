package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) block() ([]ast.Stmt, error) {
	if _, err := p.expect(token.LBRACE); err != nil {
		return nil, err
	}
	var stmts []ast.Stmt
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok == token.RBRACE {
			p.expect(token.RBRACE)
			break
		}
		stmt, err := p.stmt()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}
	return nil, nil
}

func (p *parser) stmt() (ast.Stmt, error) {
	return nil, nil
}

func (p *parser) varStmt() (*ast.VarStmt, error) {
	varName, err := p.name()
	if err != nil {
		return nil, err
	}
	varType, err := p.typ()
	if err != nil {
		return nil, err
	}
	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	if tok == token.SEMICOLON {
		p.expect(token.SEMICOLON)
		return &ast.VarStmt{
			Name: varName,
			Type: varType,
		}, nil
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
