package ast

import "github.com/nurtai325/qurt/internal/token"

type Node interface {
	aNode()
}

type node struct{}

func (*node) aNode() {}

// Declarations
// ----------------------------------------------------------------------------

type (
	Decl interface {
		Node
		aDecl()
	}

	StructDecl struct {
		Name   *NameExpr
		Fields []*Field
		decl
	}

	VarDecl struct {
		Name *NameExpr
		Type Kind
		decl
	}

	FuncDecl struct {
		Name *NameExpr
		Args []Kind
		Body []Stmt
		decl
	}
)

type decl struct {
	node
}

func (*decl) aDecl() {}

// Expressions
// ----------------------------------------------------------------------------

type (
	Expr interface {
		Node
		aExpr()
	}

	NameExpr struct {
		Value string
		expr
	}

	StringLitExpr struct {
		Value string
		expr
	}

	IntLitExpr struct {
		Value int
		expr
	}

	FloatLitExpr struct {
		Value float32
		expr
	}

	OperationExpr struct {
		Op   token.Operator
		X, Y Expr
		expr
	}

	CallExpr struct {
		Func    *NameExpr
		ArgList []Expr
		expr
	}

	SelectorExpr struct {
		Struct *NameExpr
		Field  *Field
		expr
	}

	Field struct {
		Name string
		Type Kind
	}
)

type expr struct {
	node
}

func (*expr) aExpr() {}

// Statements
// ----------------------------------------------------------------------------

type (
	Stmt interface {
		Node
		aStmt()
	}

	AssignStmt struct {
		Lhs *NameExpr
		Rhs Expr
		stmt
	}

	BreakStmt struct {
		stmt
	}

	ContinueStmt struct {
		stmt
	}

	CallStmt struct {
		CallExpr *CallExpr
		stmt
	}

	ReturnStmt struct {
		Value Expr
		stmt
	}

	IfStmt struct {
		Cond Expr
		Then []Stmt
		Else Stmt // either nil, *IfStmt, or *BlockStmt
		stmt
	}

	ForStmt struct {
		Init *VarDecl
		Cond Expr
		Post Stmt
		Body []Stmt
		stmt
	}
)

type stmt struct {
	node
}

func (*stmt) aStmt() {}

// Types
// ----------------------------------------------------------------------------

// TODO: when adding types like slice and maps make use expressions instead of just Kind enum. Move builtin types to new types package

type Kind int

const (
	StructKind Kind = iota
	IntKind
	FloatKind
)
