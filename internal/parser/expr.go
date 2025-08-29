package parser

import "github.com/nurtai325/qurtc/internal/ast"

func (p *parser) expr() (ast.Expr, error)

func (p *parser) stmt() (ast.Stmt, error)
