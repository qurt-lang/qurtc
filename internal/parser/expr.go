package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) expr() (ast.Expr, error) {
	return nil, nil
}

func (p *parser) callExpr(name *ast.NameExpr) (*ast.CallExpr, error) {
	if _, err := p.expect(token.LPAREN); err != nil {
		return nil, err
	}
	var args []ast.Expr
	for {
		arg, err := p.expr()
		if err != nil {
			return nil, err
		}
		args = append(args, arg)

		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok == token.COMMA {
			p.expect(token.COMMA)
			continue
		}
		break
	}
	if _, err := p.expect(token.RPAREN); err != nil {
		return nil, err
	}
	return &ast.CallExpr{
		Func: name,
		Args: args,
	}, nil
}
