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
	TRUE   // иә
	FALSE  // жоқ
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
	operator_end

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;

	keyword_beg
	BREAK    // тоқта
	CONTINUE // өткіз

	ELSE // әйтпесе
	FOR  // қайтала

	FUNC   // функция
	IF     // егер
	RETURN // қайтар

	STRUCT // құрылым
	VAR    // айнымалы
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ҚАТЕ",
	EOF:     "EOF",

	IDENT:  "АТАУ",
	INT:    "БҮТІН",
	FLOAT:  "БӨЛШЕК",
	STRING: "ЖОЛ",
	TRUE:   "иә",
	FALSE:  "жоқ",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	MOD: "%",

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
	SEMICOLON: ";",

	BREAK:    "тоқта",
	CONTINUE: "өткіз",

	ELSE: "әйтпесе",
	FOR:  "қайтала",

	FUNC: "функция",
	IF:   "егер",

	RETURN: "қайтар",

	STRUCT: "құрылым",
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

func (t Token) IsOperator() bool { return operator_beg < t && t < operator_end }

func (t Token) IsKeyword() bool { return keyword_beg < t && t < keyword_end }
