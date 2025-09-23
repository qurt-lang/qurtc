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
		Body       Stmts
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
		Value string
		expr
	}

	FloatLitExpr struct {
		Value string
		expr
	}

	BoolLitExpr struct {
		Value string
		expr
	}

	OperationExpr struct {
		Op   token.Token
		X, Y Expr
		expr
	}

	CallExpr struct {
		Func *NameExpr
		Args []Expr // if nil then no args
		expr
	}

	SelectorExpr struct {
		Struct Expr
		Field  *Field
		expr
	}

	ArrayAccessExpr struct {
		Array Expr
		Index Expr
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
		Var *NameExpr
		Val Expr
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
		Then Stmts
		Else Stmt // either nil, *IfStmt or Stmts
		stmt
	}

	ForStmt struct {
		Init *VarDecl
		Cond Expr
		Post Stmt
		Body Stmts
		stmt
	}

	VarDeclStmt struct {
		VarDecl *VarDecl
		stmt
	}
)

type stmt struct {
	node
}

func (*stmt) aStmt() {}

type Stmts []Stmt

func (Stmts) aStmt() {}

func (Stmts) aNode() {}

// Types
// ----------------------------------------------------------------------------

type Type struct {
	Kind     Kind
	Name     *NameExpr
	IsArray  bool
	ArrayLen int
}

type Kind int

func GetKind(typeName string) Kind {
	switch typeName {
	case types[TVoid]:
		return TVoid
	case types[TInt]:
		return TInt
	case types[TFloat]:
		return TFloat
	case types[TString]:
		return TString
	default:
		return TStruct
	}
}

func (k Kind) String() string {
	switch k {
	case TVoid:
		return types[TVoid]
	case TInt:
		return types[TInt]
	case TFloat:
		return types[TFloat]
	case TString:
		return types[TString]
	default:
		return token.STRUCT.String()
	}
}

const (
	TVoid Kind = iota
	TInt
	TFloat
	TString
	TStruct
)

var types = [...]string{
	TVoid:   "ештеңе",
	TInt:    "бүтін",
	TFloat:  "бөлшек",
	TString: "жол",
}
