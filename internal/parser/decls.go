package parser

import (
	"github.com/nurtai325/qurtc/internal/ast"
)

func (p *parser) varDecl() (ast.Decl, error) {
	varStmt, err := p.varStmt()
	if err != nil {
		return nil, err
	}
	return &ast.VarDecl{
		VarStmt: varStmt,
	}, nil
}
