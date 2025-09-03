package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) expr() (ast.Expr, error) {
	x, err := p.orTerm()
	if err != nil {
		return nil, err
	}
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok != token.LOR {
			break
		}
		p.expect(token.LOR)
		y, err := p.orTerm()
		if err != nil {
			return nil, err
		}
		x = &ast.OperationExpr{
			Op: token.LOR,
			X:  x,
			Y:  y,
		}
	}
	return x, nil
}

func (p *parser) orTerm() (ast.Expr, error) {
	x, err := p.andTerm()
	if err != nil {
		return nil, err
	}
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok != token.LAND {
			break
		}
		p.expect(token.LAND)
		y, err := p.andTerm()
		if err != nil {
			return nil, err
		}
		x = &ast.OperationExpr{
			Op: token.LAND,
			X:  x,
			Y:  y,
		}
	}
	return x, nil
}

func (p *parser) andTerm() (ast.Expr, error) {
	x, err := p.cmpTerm()
	if err != nil {
		return nil, err
	}
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		switch tok {
		case token.EQL, token.NEQ, token.LSS, token.LEQ, token.GTR, token.GEQ:
			p.expect(tok)
			y, err := p.cmpTerm()
			if err != nil {
				return nil, err
			}
			x = &ast.OperationExpr{
				Op: tok,
				X:  x,
				Y:  y,
			}
			continue
		}
		break
	}
	return x, nil
}

func (p *parser) cmpTerm() (ast.Expr, error) {
	x, err := p.addTerm()
	if err != nil {
		return nil, err
	}
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		switch tok {
		case token.ADD, token.SUB:
			p.expect(tok)
			y, err := p.addTerm()
			if err != nil {
				return nil, err
			}
			x = &ast.OperationExpr{
				Op: tok,
				X:  x,
				Y:  y,
			}
			continue
		}
		break
	}
	return x, nil
}

func (p *parser) addTerm() (ast.Expr, error) {
	x, err := p.mulTerm()
	if err != nil {
		return nil, err
	}
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		switch tok {
		case token.MUL, token.DIV, token.MOD:
			p.expect(tok)
			y, err := p.mulTerm()
			if err != nil {
				return nil, err
			}
			x = &ast.OperationExpr{
				Op: tok,
				X:  x,
				Y:  y,
			}
			continue
		}
		break
	}
	return x, nil
}

func (p *parser) mulTerm() (ast.Expr, error) {
	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	switch tok {
	case token.INT:
		lit, _ := p.expect(token.INT)
		return &ast.IntLitExpr{
			Value: lit,
		}, nil
	case token.FLOAT:
		lit, _ := p.expect(token.FLOAT)
		return &ast.FloatLitExpr{
			Value: lit,
		}, nil
	case token.STRING:
		lit, _ := p.expect(token.STRING)
		return &ast.StringLitExpr{
			Value: lit,
		}, nil
	case token.TRUE, token.FALSE:
		lit, _ := p.expect(tok)
		return &ast.BoolLitExpr{
			Value: lit,
		}, nil
	case token.IDENT:
		name, err := p.name()
		if err != nil {
			return nil, err
		}
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		switch tok {
		case token.LPAREN:
			return p.callExpr(name)
		case token.PERIOD:
			return p.selectorExpr(name)
		case token.LBRACK:
			return p.arrayAccessExpr(name)
		default:
			return name, nil
		}
	default:
		return nil, ErrUnexpectedToken
	}
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

func (p *parser) selectorExpr(structExpr ast.Expr) (ast.Expr, error) {
	if _, err := p.expect(token.PERIOD); err != nil {
		return nil, err
	}
	field, err := p.name()
	if err != nil {
		return nil, err
	}
	return &ast.SelectorExpr{
		Struct: structExpr,
		Field: &ast.Field{
			Name: field.Value,
		},
	}, nil
}

func (p *parser) arrayAccessExpr(arrayExpr ast.Expr) (ast.Expr, error) {
	if _, err := p.expect(token.LBRACK); err != nil {
		return nil, err
	}
	index, err := p.expr()
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(token.RBRACK); err != nil {
		return nil, err
	}
	return &ast.ArrayAccessExpr{
		Array: arrayExpr,
		Index: index,
	}, nil
}
