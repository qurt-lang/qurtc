package parser

import (
	"errors"
	"strconv"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/token"
)

type precedence int

const (
	_ precedence = iota
	precOrOr
	precAndAnd
	precCmp
	precAdd
	precMul
	precUnary
)

var precs = map[token.Token]precedence{
	token.MUL: precMul,
	token.DIV: precMul,
	token.MOD: precMul,

	token.ADD: precAdd,
	token.SUB: precAdd,

	token.EQL: precCmp,
	token.LSS: precCmp,
	token.GTR: precCmp,
	token.NEQ: precCmp,
	token.LEQ: precCmp,
	token.GEQ: precCmp,

	token.LAND: precAndAnd,
	token.LOR:  precOrOr,
}

func (p *parser) expr(prec precedence) (ast.Expr, error) {
	tok, err := p.peek()
	if err != nil {
		return nil, errors.Join(ErrInvalidExpr, err)
	}
	prefix := p.prefixFuncs[tok]
	if prefix == nil {
		return nil, ErrInvalidExpr
	}
	left, err := prefix()
	if err != nil {
		return nil, errors.Join(ErrInvalidExpr, err)
	}
	nextTok, err := p.peek()
	if err != nil {
		return nil, errors.Join(ErrInvalidExpr, err)
	}
	for nextTok != token.SEMICOLON && prec < p.prec(nextTok) {
		infix := p.infixFuncs[nextTok]
		if infix == nil {
			return left, nil
		}
		left, err = infix(left)
		if err != nil {
			return nil, errors.Join(ErrInvalidExpr, err)
		}
		nextTok, err = p.peek()
		if err != nil {
			return nil, errors.Join(ErrInvalidExpr, err)
		}
	}
	return left, nil
}

func (p *parser) prefix() (ast.Expr, error) {
	_, err := p.expect(token.SUB, token.NOT)
	if err != nil {
		return nil, err
	}
	op := p.s.Tok()
	operand, err := p.expr(precUnary)
	if err != nil {
		return nil, err
	}
	return &ast.UnaryOpExpr{
		Operand: operand,
		Op:      op,
	}, nil
}

func (p *parser) infix(left ast.Expr) (ast.Expr, error) {
	tok, err := p.peek()
	if err != nil {
		return nil, err
	}
	prec := p.prec(tok)
	expr := ast.OpExpr{
		Left: left,
		Op:   tok,
	}
	p.expect(tok)

	right, err := p.expr(prec)
	if err != nil {
		return nil, err
	}
	expr.Right = right
	return &expr, nil
}

func (p *parser) prec(op token.Token) precedence {
	if currPrec, ok := precs[op]; ok {
		return currPrec
	}
	return 0
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
		_, err := p.expect(token.PERIOD)
		if err != nil {
			return nil, err
		}

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
	index, err := p.expr(0)
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
		_, err := p.expect(token.TRUE)
		if err != nil {
			return nil, errors.Join(ErrInvalidBool, err)
		}
		return &ast.BoolExpr{Value: true}, nil
	} else if tok == token.FALSE {
		_, err := p.expect(token.FALSE)
		if err != nil {
			return nil, errors.Join(ErrInvalidBool, err)
		}
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

		element, err := p.expr(0)
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
		_, err = p.expect(token.COMMA)
		if err != nil {
			return nil, err
		}
	}
	_, err := p.expect(end)
	if err != nil {
		return nil, err
	}
	return exprs, nil
}
