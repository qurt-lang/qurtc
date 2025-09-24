package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/help"
	"github.com/nurtai325/qurtc/internal/scanner"
	"github.com/nurtai325/qurtc/internal/token"
)

type parser struct {
	s           scanner.Scanner
	prefixFuncs map[token.Token]func() (ast.Expr, error)
	infixFuncs  map[token.Token]func(left ast.Expr) (ast.Expr, error)
}

func New(filename string, input []byte) *parser {
	newParser := &parser{
		s: scanner.New(filename, []byte(input)),
	}
	newParser.prefixFuncs = map[token.Token]func() (ast.Expr, error){
		token.IDENT:  newParser.nameExpr,
		token.LBRACE: newParser.array,
		token.STRING: newParser.string,
		token.INT:    newParser.int,
		token.FLOAT:  newParser.float,
		token.TRUE:   newParser.bool,
		token.FALSE:  newParser.bool,
	}
	return newParser
}

func (p *parser) Parse() (decls []ast.Decl, err error) {
	// TODO: add known errors, return them from functions without panicking
	for p.s.Scan() {
		switch p.s.Tok() {
		// case token.FUNC:
		// 	decls, err = p.appendDecl(decls, p.funcDecl)
		// 	if err != nil {
		// 		return nil, p.errorAt(errors.Join(ErrInvalidFuncDecl, err), help.FunctionsPage)
		// 	}
		// case token.STRUCT:
		// 	decls, err = p.appendDecl(decls, p.structDecl)
		// 	if err != nil {
		// 		return nil, p.errorAt(errors.Join(ErrInvalidStructDecl, err), help.StructsPage)
		// 	}
		case token.VAR:
			decls, err = p.appendDecl(decls, p.varDecl)
			if err != nil {
				return nil, p.errorAt(errors.Join(ErrInvalidVarDecl, err), help.VarsPage)
			}
		default:
			return nil, p.errorAt(ErrUnknownDecl, help.SyntaxPage)
		}
	}
	if p.s.Tok() == token.ILLEGAL {
		return nil, p.errorAt(p.s.Err(), help.QurtTour)
	}
	return decls, nil
}

func (p *parser) appendDecl(decls []ast.Decl, fn func() (ast.Decl, error)) ([]ast.Decl, error) {
	decl, err := fn()
	if err != nil {
		return nil, err
	}
	return append(decls, decl), nil
}

func (p *parser) name() (*ast.NameExpr, error) {
	lit, err := p.expect(token.IDENT)
	if err != nil {
		return nil, errors.Join(ErrInvalidIdent, err)
	}
	return &ast.NameExpr{Value: lit}, nil
}

func (p *parser) typ() (*ast.Type, error) {
	var t ast.Type
	if tok, _ := p.peek(); tok == token.LBRACK {
		arrLen, err := p.arrlen()
		if err != nil {
			return nil, errors.Join(ErrInvalidArrayLen, err)
		}
		t.IsArray = true
		t.ArrayLen = arrLen
	}
	name, err := p.name()
	if err != nil {
		return nil, errors.Join(ErrInvalidTypeName, err)
	}
	t.Name = name
	t.Kind = ast.GetKind(name.Value)
	return &t, nil
}

func (p *parser) arrlen() (int, error) {
	p.expect(token.LBRACK)
	lit, err := p.expect(token.INT)
	if err != nil {
		return 0, err
	}
	if _, err := p.expect(token.RBRACK); err != nil {
		return 0, err
	}
	arrayLen, err := strconv.ParseInt(lit, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(arrayLen), nil
}

func (p *parser) expect(tok token.Token) (string, error) {
	if !p.s.Scan() {
		if p.s.Tok() == token.EOF {
			return "", ErrUnexpectedEOF
		} else {
			return "", p.s.Err()
		}
	} else if p.s.Tok() != tok {
		return "", fmt.Errorf("күтпеген таңба немесе cөз: %q керек, бірақ %q табылды", tok.String(), p.s.Tok().String())
	}
	return p.s.Lit(), nil
}

func (p *parser) peek() (token.Token, error) {
	tok, err := p.s.Peek()
	if tok == token.EOF {
		return token.EOF, ErrUnexpectedEOF
	}
	return tok, err
}
