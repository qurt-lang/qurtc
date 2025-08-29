package ast

import "github.com/nurtai325/qurtc/internal/token"

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
		Type *Type
		Val  Expr
		decl
	}

	FuncDecl struct {
		Name       *NameExpr
		Args       []*FuncArg
		ReturnType *Type
		Body       []Stmt
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

	ArrayExpr struct {
		Length   int
		Elements []Expr
		expr
	}

	ArrayAccessExpr struct {
		Array *NameExpr
		Index *Expr
		expr
	}

	FuncArg struct {
		Name string
		Type *Type
	}

	Field struct {
		Name string
		Type *Type
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

type Type struct {
	Kind    Kind
	Name    *NameExpr
	IsArray bool
	Length  int
}

type Kind int

const (
	Unknown Kind = iota
	TVoid
	TStruct
	TInt
	TFloat
	TString
)
