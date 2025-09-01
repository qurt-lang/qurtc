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
		name, err := p.name()
		if err != nil {
			return nil, err
		}
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok == token.ASSIGN {
			p.expect(token.ASSIGN)
			val, err := p.expr()
			if err != nil {
				return nil, err
			}
			_, err = p.expect(token.SEMICOLON)
			if err != nil {
				return nil, err
			}
			return &ast.AssignStmt{
				Var: name,
				Val: val,
			}, nil
		} else if tok == token.LPAREN {
			p.expect(token.LPAREN)
			var args []ast.Expr
			for {
				tok, err := p.peek()
				if err != nil {
					return nil, err
				} else if tok == token.RPAREN {
					break
				}
				arg, err := p.expr()
				if err != nil {
					return nil, err
				}
				_, err = p.expect(token.COMMA)
				if err != nil {
					return nil, err
				}
				args = append(args, arg)
			}
			_, err = p.expect(token.RBRACE)
			if err != nil {
				return nil, err
			}
			_, err = p.expect(token.SEMICOLON)
			if err != nil {
				return nil, err
			}
			return &ast.CallStmt{
				CallExpr: &ast.CallExpr{
					Func: name,
					Args: args,
				},
			}, nil
		} else {
			return nil, ErrUnexpectedToken
		}
	default:
		return nil, ErrUnexpectedToken
	}
}
