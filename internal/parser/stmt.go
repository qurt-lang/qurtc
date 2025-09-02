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
		return p.assignOrCallStmt()
	case token.FOR:
		return p.forStmt()
	case token.IF:
		return p.ifStmt()
	case token.VAR:
		p.expect(token.VAR)
		varDecl, err := p.varDecl()
		if err != nil {
			return nil, err
		}
		return &ast.VarDeclStmt{
			// TODO: remove this type assertion when all declarations are allowed in func body and add ast.DeclStmt
			VarDecl: varDecl.(*ast.VarDecl),
		}, nil
	case token.RETURN:
		p.expect(token.RETURN)
		returnVal, err := p.expr()
		if err != nil {
			return nil, err
		}
		_, err = p.expect(token.SEMICOLON)
		if err != nil {
			return nil, err
		}
		return &ast.ReturnStmt{
			Value: returnVal,
		}, nil
	case token.BREAK:
		p.expect(token.BREAK)
		_, err = p.expect(token.SEMICOLON)
		if err != nil {
			return nil, err
		}
		return &ast.BreakStmt{}, nil
	case token.CONTINUE:
		p.expect(token.CONTINUE)
		_, err = p.expect(token.SEMICOLON)
		if err != nil {
			return nil, err
		}
		return &ast.ContinueStmt{}, nil
	default:
		return nil, ErrUnexpectedToken
	}
}

func (p *parser) block() (ast.Stmts, error) {
	_, err := p.expect(token.LBRACE)
	if err != nil {
		return nil, err
	}
	var block ast.Stmts
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		} else if tok == token.RBRACE {
			break
		}
		stmt, err := p.stmt()
		if err != nil {
			return nil, err
		}
		block = append(block, stmt)
	}
	_, err = p.expect(token.RBRACE)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (p *parser) forStmt() (ast.Stmt, error) {
	_, err := p.expect(token.FOR)
	if err != nil {
		return nil, err
	}

	_, err = p.expect(token.LPAREN)
	if err != nil {
		return nil, err
	}
	init, err := p.varDecl()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.SEMICOLON)
	if err != nil {
		return nil, err
	}

	cond, err := p.expr()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.SEMICOLON)
	if err != nil {
		return nil, err
	}

	name, err := p.name()
	if err != nil {
		return nil, err
	}
	post, err := p.assignStmt(name)
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
		Init: init.(*ast.VarDecl),
		Cond: cond,
		Post: post,
		Body: body,
	}, nil
}

func (p *parser) ifStmt() (ast.Stmt, error) {
	_, err := p.expect(token.IF)
	if err != nil {
		return nil, err
	}

	_, err = p.expect(token.LPAREN)
	if err != nil {
		return nil, err
	}
	cond, err := p.expr()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.RPAREN)
	if err != nil {
		return nil, err
	}

	block, err := p.block()
	if err != nil {
		return nil, err
	}

	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	if tok == token.ELSE {
		p.expect(token.ELSE)
		tok, err = p.peek()
		if err != nil {
			return nil, err
		}
		if tok == token.IF {
			elseIf, err := p.ifStmt()
			if err != nil {
				return nil, err
			}
			return &ast.IfStmt{
				Cond: cond,
				Then: block,
				Else: elseIf,
			}, nil
		}
		elseBlock, err := p.block()
		if err != nil {
			return nil, err
		}
		return &ast.IfStmt{
			Cond: cond,
			Then: block,
			Else: elseBlock,
		}, nil
	}

	return &ast.IfStmt{
		Cond: cond,
		Then: block,
	}, nil
}

func (p *parser) assignOrCallStmt() (ast.Stmt, error) {
	name, err := p.name()
	if err != nil {
		return nil, err
	}
	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	switch tok {
	case token.ASSIGN:
		return p.assignStmt(name)
	case token.LPAREN:
		return p.callStmt(name)
	default:
		return nil, ErrUnexpectedToken
	}
}

func (p *parser) assignStmt(name *ast.NameExpr) (ast.Stmt, error) {
	if _, err := p.expect(token.ASSIGN); err != nil {
		return nil, err
	}
	val, err := p.expr()
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, err
	}
	return &ast.AssignStmt{
		Var: name,
		Val: val,
	}, nil
}

func (p *parser) callStmt(name *ast.NameExpr) (ast.Stmt, error) {
	callExpr, err := p.callExpr(name)
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, err
	}
	return &ast.CallStmt{
		CallExpr: callExpr,
	}, nil
}
