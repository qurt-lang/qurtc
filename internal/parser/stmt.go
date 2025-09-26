package parser

import (
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
		stmt, err := p.assignStmt()
		if err != nil {
			return nil, err
		}
		if _, err = p.expect(token.SEMICOLON); err != nil {
			return nil, err
		}
		return stmt, nil
	case token.VAR:
		p.expect(token.VAR)
		return p.varStmt()
	case token.IF:
		p.expect(token.IF)
		return p.ifStmt()
	case token.FOR:
		p.expect(token.FOR)
		return p.forStmt()
	case token.CONTINUE:
		p.expect(token.CONTINUE)
		_, err = p.expect(token.SEMICOLON)
		if err != nil {
			return nil, err
		}
		return &ast.ContinueStmt{}, nil
	case token.BREAK:
		p.expect(token.BREAK)
		_, err = p.expect(token.SEMICOLON)
		if err != nil {
			return nil, err
		}
		return &ast.BreakStmt{}, nil
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
	default:
		return nil, ErrUnknownStmt
	}
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

func (p *parser) ifStmt() (*ast.IfStmt, error) {
	var stmt *ast.IfStmt = &ast.IfStmt{}

	_, err := p.expect(token.LPAREN)
	if err != nil {
		return nil, err
	}
	stmt.Cond, err = p.expr(0)
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.RPAREN)
	if err != nil {
		return nil, err
	}
	stmt.Then, err = p.block()
	if err != nil {
		return nil, err
	}

	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	if tok == token.ELSE {
		p.expect(token.ELSE)
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok == token.IF {
			p.expect()
			stmt.Else, err = p.ifStmt()
			if err != nil {
				return nil, err
			}
		} else {
			stmtElse, err := p.block()
			if err != nil {
				return nil, err
			}
			stmt.Else = ast.Stmts(stmtElse)
		}
	}
	return stmt, nil
}

func (p *parser) forStmt() (*ast.ForStmt, error) {
	_, err := p.expect(token.LPAREN)
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.VAR)
	if err != nil {
		return nil, err
	}
	init, err := p.varStmt()
	if err != nil {
		return nil, err
	}
	cond, err := p.expr(0)
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.SEMICOLON)
	if err != nil {
		return nil, err
	}
	post, err := p.assignStmt()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.RPAREN)
	if err != nil {
		return nil, err
	}
	body, err := p.block()
	if err != nil {
		return nil, err
	}
	return &ast.ForStmt{
		Init: init,
		Cond: cond,
		Post: post,
		Body: body,
	}, nil
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

func (p *parser) assignStmt() (ast.Stmt, error) {
	assignee, err := p.nameExpr()
	if err != nil {
		return nil, err
	}
	if call, ok := assignee.(*ast.CallExpr); ok {
		return &ast.CallStmt{
			CallExpr: call,
		}, nil
	}
	p.expect(token.ASSIGN)
	val, err := p.expr(0)
	if err != nil {
		return nil, err
	}
	return &ast.AssignStmt{
		Var: assignee,
		Val: val,
	}, nil
}
