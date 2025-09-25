package parser

import (
	"fmt"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) stmt() (ast.Stmt, error) {
	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	switch tok {
	case token.IDENT:
		varName, err := p.name()
		if err != nil {
			return nil, err
		}
		return p.assignOrCall(varName)
	case token.VAR:
		p.expect(token.VAR)
		return p.varStmt()
	case token.RETURN:
		p.expect(token.RETURN)
		val, err := p.expr(0)
		if err != nil {
			return nil, err
		}
		_, err = p.expect(token.SEMICOLON)
		if err != nil {
			return nil, err
		}
		return &ast.ReturnStmt{
			Value: val,
		}, nil
	}
	return nil, nil
}

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
	return stmts, nil
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

func (p *parser) assignOrCall(varName *ast.NameExpr) (ast.Stmt, error) {
	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	switch tok {
	case token.PERIOD, token.LBRACK, token.ASSIGN:
		var assignee ast.Expr
		switch tok {
		case token.PERIOD:
			assignee, err = p.selector(varName)
			if err != nil {
				return nil, err
			}
		case token.LBRACK:
			assignee, err = p.arrayAccess(varName)
			if err != nil {
				return nil, err
			}
		case token.ASSIGN:
			assignee = varName
		default:
			return nil, ErrUnknownStmt
		}

		p.expect(token.ASSIGN)
		val, err := p.expr(0)
		if err != nil {
			return nil, err
		}
		if _, err = p.expect(token.SEMICOLON); err != nil {
			return nil, err
		}
		return &ast.AssignStmt{
			Var: assignee,
			Val: val,
		}, nil
	case token.LPAREN:
		call, err := p.call(varName)
		if err != nil {
			return nil, err
		}
		return &ast.CallStmt{
			CallExpr: call,
		}, nil
	default:
		fmt.Println(varName.Value, tok)
		return nil, ErrUnknownStmt
	}
}
