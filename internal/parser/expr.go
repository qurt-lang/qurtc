package parser

import (
	"errors"
	"strconv"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

func (p *parser) expr() (ast.Expr, error) {
	tok, err := p.peek()
	if err != nil {
		return nil, errors.Join(ErrInvalidExpr, err)
	}
	fn, ok := p.prefixFuncs[tok]
	if !ok {
		return nil, nil
	}
	return fn()
}

func (p *parser) infix() (ast.Expr, error) {
	return nil, nil
}

func (p *parser) prefix(left ast.Expr) (ast.Expr, error) {
	return nil, nil
}

func (p *parser) nameExpr() (ast.Expr, error) {
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
		return p.call(name)
	case token.PERIOD:
		return p.selector(name)
	case token.LBRACK:
		return p.arrayAccess(name)
	default:
		return name, nil
	}
}

func (p *parser) call(name *ast.NameExpr) (*ast.CallExpr, error) {
	_, err := p.expect(token.LPAREN)
	if err != nil {
		return nil, errors.Join(ErrInvalidFuncCall, err)
	}
	args, err := p.exprList(token.RPAREN)
	if err != nil {
		return nil, errors.Join(ErrInvalidFuncCall, err)
	}
	return &ast.CallExpr{
		Func: name,
		Args: args,
	}, nil
}

func (p *parser) selector(name *ast.NameExpr) (*ast.SelectorExpr, error) {
	selectorExpr := &ast.SelectorExpr{
		Field: name,
	}
	for {
		p.expect(token.PERIOD)

		name, err := p.name()
		if err != nil {
			return nil, err
		}
		selectorExpr = &ast.SelectorExpr{
			Struct: selectorExpr,
			Field:  name,
		}

		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok != token.PERIOD {
			break
		}
	}
	return selectorExpr, nil
}

func (p *parser) arrayAccess(name *ast.NameExpr) (*ast.ArrayAccessExpr, error) {
	_, err := p.expect(token.LBRACK)
	if err != nil {
		return nil, err
	}
	index, err := p.expr()
	if err != nil {
		return nil, err
	}
	_, err = p.expect(token.RBRACK)
	if err != nil {
		return nil, err
	}
	return &ast.ArrayAccessExpr{
		Array: name,
		Index: index,
	}, nil
}

func (p *parser) array() (ast.Expr, error) {
	_, err := p.expect(token.LBRACE)
	if err != nil {
		return nil, errors.Join(ErrInvalidArray, err)
	}
	elements, err := p.exprList(token.RBRACE)
	if err != nil {
		return nil, errors.Join(ErrInvalidArray, err)
	}
	return &ast.ArrayExpr{
		Elements: elements,
	}, nil
}

func (p *parser) string() (ast.Expr, error) {
	lit, err := p.expect(token.STRING)
	if err != nil {
		return nil, errors.Join(ErrInvalidString, err)
	}
	return &ast.StringExpr{Value: lit}, nil
}

func (p *parser) int() (ast.Expr, error) {
	lit, err := p.expect(token.INT)
	if err != nil {
		return nil, errors.Join(ErrInvalidInt, err)
	}
	val, err := strconv.ParseInt(lit, 10, 0)
	if err != nil {
		return nil, errors.Join(ErrInvalidInt, err)
	}
	return &ast.IntExpr{Value: int(val)}, nil
}

func (p *parser) float() (ast.Expr, error) {
	lit, err := p.expect(token.FLOAT)
	if err != nil {
		return nil, errors.Join(ErrInvalidFloat, err)
	}
	val, err := strconv.ParseFloat(lit, 32)
	if err != nil {
		return nil, errors.Join(ErrInvalidFloat, err)
	}
	return &ast.FloatExpr{Value: float32(val)}, nil
}

func (p *parser) bool() (ast.Expr, error) {
	tok, err := p.peek()
	if err != nil {
		return nil, errors.Join(ErrInvalidBool, err)
	}
	if tok == token.TRUE {
		p.expect(token.TRUE)
		return &ast.BoolExpr{Value: true}, nil
	} else if tok == token.FALSE {
		p.expect(token.FALSE)
		return &ast.BoolExpr{Value: false}, nil
	}
	return nil, errors.Join(ErrInvalidBool, err)
}

func (p *parser) exprList(end token.Token) ([]ast.Expr, error) {
	var exprs []ast.Expr
	for {
		tok, err := p.peek()
		if err != nil {
			return nil, err
		}
		if tok == end {
			break
		}

		element, err := p.expr()
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, element)

		tok, err = p.peek()
		if err != nil {
			return nil, err
		}
		if tok != token.COMMA {
			break
		}
		p.expect(token.COMMA)
	}
	p.expect(end)
	return exprs, nil
}
