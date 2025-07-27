package token

type Token int

const (
	ILLEGAL Token = iota
	EOF

	literal_beg
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	STRING // "abc"
	literal_end

	operator_beg
	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %

	LAND // &&
	LOR  // ||

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ // !=
	LEQ // <=
	GEQ // >=

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	COLON     // :
	SEMICOLON // ;
	operator_end

	keyword_beg
	BREAK
	CASE
	CONTINUE

	DEFAULT
	ELSE
	FOR

	FUNC
	IF
	RETURN

	STRUCT
	SWITCH
	VAR
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ҚАТЕ",

	EOF: "EOF",

	IDENT:  "АТАУ",
	INT:    "БҮТІН",
	FLOAT:  "БӨЛШЕК",
	STRING: "МӘТІН",

	ADD: "+",
	SUB: "-",
	MUL: "*",

	LAND: "&&",
	LOR:  "||",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ: "!=",
	LEQ: "<=",
	GEQ: ">=",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	COLON:     ":",
	SEMICOLON: ";",

	BREAK:    "тоқта",
	CASE:     "нұсқа",
	CONTINUE: "өткіз",

	DEFAULT: "әдепкі",
	ELSE:    "әйтпесе",
	FOR:     "қайтала",

	FUNC: "функция",
	IF:   "егер",

	RETURN: "қайтар",

	STRUCT: "құрылым",
	SWITCH: "таңда",
	VAR:    "айнымалы",
}

func (t Token) String() string {
	return tokens[t]
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(ident string) (Token, bool) {
	if tok, ok := keywords[ident]; ok {
		return tok, true
	}
	return 0, false
}

func (t Token) IsLiteral() bool { return literal_beg < t && t < literal_end }

func (t Token) IsOperator() bool { return literal_beg < t && t < literal_end }

func (t Token) IsKeyword() bool { return literal_beg < t && t < literal_end }
