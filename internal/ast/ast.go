package ast

import (
	"github.com/nurtai325/qurtc/internal/token"
)

// Declarations
// ----------------------------------------------------------------------------

type (
	Decl interface {
		aDecl()
	}

	StructDecl struct {
		Name   *NameExpr
		Fields []*Field
		decl
	}

	FuncDecl struct {
		Name       *NameExpr
		Args       []*FuncArg
		ReturnType *Type
		Body       []Stmt
		decl
	}

	// only for builtin funcs
	BuiltinFuncDecl struct {
		Name *NameExpr
		Body func(...any) error
		decl
	}
)

type decl struct {
}

func (*decl) aDecl() {}

type FuncArg struct {
	Name string
	Type *Type
}

type Field struct {
	Name string
	Type *Type
}

// Expressions
// ----------------------------------------------------------------------------

type (
	Expr interface {
		aExpr()
	}

	NameExpr struct {
		Value string
		expr
	}

	StringExpr struct {
		Value string
		expr
	}

	IntExpr struct {
		Value int
		expr
	}

	FloatExpr struct {
		Value float32
		expr
	}

	BoolExpr struct {
		Value bool
		expr
	}

	ArrayExpr struct {
		Elements []Expr
		expr
	}

	CallExpr struct {
		Func Expr
		Args []Expr // if nil then no args
		expr
	}

	SelectorExpr struct {
		Struct Expr
		Field  *NameExpr
		expr
	}

	ArrayAccessExpr struct {
		Array Expr
		Index Expr
		expr
	}

	UnaryOpExpr struct {
		Op      token.Token
		Operand Expr
		expr
	}

	OpExpr struct {
		Op          token.Token
		Left, Right Expr
		expr
	}
)

type expr struct {
}

func (*expr) aExpr() {}

// Statements
// ----------------------------------------------------------------------------

type (
	Stmt interface {
		aStmt()
	}

	VarStmt struct {
		Name *NameExpr
		Type *Type
		Val  Expr // zero value if nil
		stmt
	}

	AssignStmt struct {
		Var Expr // *ArrayAccess, *Selector, *NameExpr
		Val Expr
		stmt
	}

	CallStmt struct {
		CallExpr *CallExpr
		stmt
	}

	IfStmt struct {
		Cond Expr
		Then []Stmt
		Else Stmt // either nil, *IfStmt or Stmts
		stmt
	}

	ForStmt struct {
		Init *VarStmt
		Cond Expr
		Post Stmt
		Body []Stmt
		stmt
	}

	ReturnStmt struct {
		Value Expr
		stmt
	}

	BreakStmt struct {
		stmt
	}

	ContinueStmt struct {
		stmt
	}
)

type stmt struct {
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
	for i, kindName := range types {
		if typeName == kindName {
			return Kind(i)
		}
	}
	return TStruct
}

func (k Kind) String() string {
	if int(k) >= len(types) {
		return token.STRUCT.String()
	}
	return types[k]
}

const (
	TVoid Kind = iota
	TInt
	TFloat
	TString
	TBool
	TStruct
)

var types = [...]string{
	TVoid:   "ештеңе",
	TInt:    "бүтін",
	TFloat:  "бөлшек",
	TString: "жол",
	TBool:   "шын",
}
