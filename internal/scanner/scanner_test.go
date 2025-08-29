package scanner_test

import (
	"strings"
	"testing"

	"github.com/nurtai325/qurtc/internal/scanner"
	"github.com/nurtai325/qurtc/internal/token"
)

type scannerTestCase struct {
	tok token.Token
	lit string
}

type scannerTest struct {
	name   string
	input  string
	tokens []scannerTestCase
}

var tests []scannerTest = []scannerTest{
	{
		name:  "Empty input",
		input: "",
		tokens: []scannerTestCase{
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Whitespace only",
		input: "   \t\t\n  ",
		tokens: []scannerTestCase{
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Test IsLiteral method",
		input: "x 123 45.67 \"text\"",
		tokens: []scannerTestCase{
			{token.IDENT, "x"},
			{token.INT, "123"},
			{token.FLOAT, "45.67"},
			{token.STRING, "text"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Test IsOperator method",
		input: "+ - * / % && || == != <= >= < > ! = ( ) [ ] { } , . : ;",
		tokens: []scannerTestCase{
			{token.ADD, "+"}, {token.SUB, "-"}, {token.MUL, "*"}, {token.DIV, "/"}, {token.MOD, "%"},
			{token.LAND, "&&"}, {token.LOR, "||"}, {token.EQL, "=="}, {token.NEQ, "!="},
			{token.LEQ, "<="}, {token.GEQ, ">="}, {token.LSS, "<"}, {token.GTR, ">"},
			{token.NOT, "!"}, {token.ASSIGN, "="}, {token.LPAREN, "("}, {token.RPAREN, ")"},
			{token.LBRACK, "["}, {token.RBRACK, "]"}, {token.LBRACE, "{"}, {token.RBRACE, "}"},
			{token.COMMA, ","}, {token.PERIOD, "."}, {token.COLON, ":"}, {token.SEMICOLON, "semicolon"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Test IsKeyword method",
		input: "тоқта өткіз әйтпесе қайтала функция егер қайтар құрылым айнымалы",
		tokens: []scannerTestCase{
			{token.BREAK, "тоқта"}, {token.CONTINUE, "өткіз"}, {token.ELSE, "әйтпесе"},
			{token.FOR, "қайтала"}, {token.FUNC, "функция"}, {token.IF, "егер"},
			{token.RETURN, "қайтар"}, {token.STRUCT, "құрылым"}, {token.VAR, "айнымалы"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Newlines treated as whitespace",
		input: "егер(x==5){\nқайтар 10\n}\n",
		tokens: []scannerTestCase{
			{token.IF, "егер"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.EQL, "=="},
			{token.INT, "5"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RETURN, "қайтар"},
			{token.INT, "10"},
			{token.RBRACE, "}"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Mixed line endings",
		input: "x\ny\nz\t",
		tokens: []scannerTestCase{
			{token.IDENT, "x"},
			{token.IDENT, "y"},
			{token.IDENT, "z"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Integer edge cases",
		input: "0 00 123 1000000000000000000000",
		tokens: []scannerTestCase{
			{token.INT, "0"},
			{token.INT, "00"},
			{token.INT, "123"},
			{token.INT, "1000000000000000000000"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Float edge cases",
		input: "0.0 .5 123. 0.123 999.999",
		tokens: []scannerTestCase{
			{token.FLOAT, "0.0"},
			{token.PERIOD, "."},
			{token.INT, "5"},
			{token.FLOAT, "123."},
			{token.FLOAT, "0.123"},
			{token.FLOAT, "999.999"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Malformed float variations",
		input: "12..45 123...456 .",
		tokens: []scannerTestCase{
			{token.FLOAT, "12."},
			{token.PERIOD, "."},
			{token.INT, "45"},
			{token.FLOAT, "123."},
			{token.PERIOD, "."},
			{token.PERIOD, "."},
			{token.INT, "456"},
			{token.PERIOD, "."},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Numbers followed by identifiers",
		input: "123abc 45.67def 0x123",
		tokens: []scannerTestCase{
			{token.INT, "123"},
			{token.IDENT, "abc"},
			{token.FLOAT, "45.67"},
			{token.IDENT, "def"},
			{token.INT, "0"},
			{token.IDENT, "x123"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "String variations",
		input: "\"hello\" \"\" \"қазақша мәтін\" \"with spaces\"",
		tokens: []scannerTestCase{
			{token.STRING, "hello"},
			{token.STRING, ""},
			{token.STRING, "қазақша мәтін"},
			{token.STRING, "with spaces"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "String with escape sequences",
		input: "\"line1\nline2\" \"tab\there\" \"quote\\\"inside\"",
		tokens: []scannerTestCase{
			{token.STRING, "line1\nline2"},
			{token.STRING, "tab\there"},
			{token.STRING, "quote\\\"inside"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "String with newlines",
		input: "\"line1\nline2\"",
		tokens: []scannerTestCase{
			{token.STRING, "line1\nline2"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Unicode identifiers",
		input: "айнымалы сөзҰзын123 mixedАБВ кириллица",
		tokens: []scannerTestCase{
			{token.VAR, "айнымалы"},
			{token.IDENT, "сөзҰзын123"},
			{token.IDENT, "mixedАБВ"},
			{token.IDENT, "кириллица"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Keywords vs identifiers",
		input: "егер егерсіз функцияFoo тоқтаBar",
		tokens: []scannerTestCase{
			{token.IF, "егер"},
			{token.IDENT, "егерсіз"},
			{token.IDENT, "функцияFoo"},
			{token.IDENT, "тоқтаBar"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Operator precedence order",
		input: "!x&&y||z==w!=v<=u>=t<s>r",
		tokens: []scannerTestCase{
			{token.NOT, "!"}, {token.IDENT, "x"}, {token.LAND, "&&"}, {token.IDENT, "y"},
			{token.LOR, "||"}, {token.IDENT, "z"}, {token.EQL, "=="}, {token.IDENT, "w"},
			{token.NEQ, "!="}, {token.IDENT, "v"}, {token.LEQ, "<="}, {token.IDENT, "u"},
			{token.GEQ, ">="}, {token.IDENT, "t"}, {token.LSS, "<"}, {token.IDENT, "s"},
			{token.GTR, ">"}, {token.IDENT, "r"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Ambiguous operator sequences",
		input: "<= >= == != ! = < > && || & ",
		tokens: []scannerTestCase{
			{token.LEQ, "<="}, {token.GEQ, ">="}, {token.EQL, "=="}, {token.NEQ, "!="},
			{token.NOT, "!"}, {token.ASSIGN, "="}, {token.LSS, "<"}, {token.GTR, ">"},
			{token.LAND, "&&"}, {token.LOR, "||"}, {token.ILLEGAL, "ҚАТЕ"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Single illegal character",
		input: `#`,
		tokens: []scannerTestCase{
			{token.ILLEGAL, "ҚАТЕ"},
		},
	},
	{
		name:  "Valid token followed by illegal character",
		input: `айнымалы @`,
		tokens: []scannerTestCase{
			{token.VAR, "айнымалы"},
			{token.ILLEGAL, "ҚАТЕ"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Function definition with complex body",
		input: "функция factorial(n) {\nегер (n <= 1) {\nқайтар 1;\n} әйтпесе {\nқайтар n * factorial(n - 1);}}",
		tokens: []scannerTestCase{
			{token.FUNC, "функция"}, {token.IDENT, "factorial"}, {token.LPAREN, "("},
			{token.IDENT, "n"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
			{token.IF, "егер"}, {token.LPAREN, "("}, {token.IDENT, "n"}, {token.LEQ, "<="},
			{token.INT, "1"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
			{token.RETURN, "қайтар"}, {token.INT, "1"}, {token.SEMICOLON, "semicolon"},
			{token.RBRACE, "}"}, {token.ELSE, "әйтпесе"}, {token.LBRACE, "{"},
			{token.RETURN, "қайтар"}, {token.IDENT, "n"}, {token.MUL, "*"}, {token.IDENT, "factorial"},
			{token.LPAREN, "("}, {token.IDENT, "n"}, {token.SUB, "-"}, {token.INT, "1"},
			{token.RPAREN, ")"}, {token.SEMICOLON, "semicolon"},
			{token.RBRACE, "}"}, {token.RBRACE, "}"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Struct definition",
		input: "құрылым Point {\nx: INT,\ny: INT\n}",
		tokens: []scannerTestCase{
			{token.STRUCT, "құрылым"}, {token.IDENT, "Point"}, {token.LBRACE, "{"},
			{token.IDENT, "x"}, {token.COLON, ":"}, {token.IDENT, "INT"}, {token.COMMA, ","},
			{token.IDENT, "y"}, {token.COLON, ":"}, {token.IDENT, "INT"},
			{token.RBRACE, "}"},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Loop with break and continue",
		input: "қайтала {\nегер (x == 0) тоқта;\nегер (x % 2 == 0) өткіз;\nx = x - 1;\n}",
		tokens: []scannerTestCase{
			{token.FOR, "қайтала"}, {token.LBRACE, "{"},
			{token.IF, "егер"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.EQL, "=="},
			{token.INT, "0"}, {token.RPAREN, ")"}, {token.BREAK, "тоқта"}, {token.SEMICOLON, "semicolon"},
			{token.IF, "егер"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.MOD, "%"},
			{token.INT, "2"}, {token.EQL, "=="}, {token.INT, "0"}, {token.RPAREN, ")"},
			{token.CONTINUE, "өткіз"}, {token.SEMICOLON, "semicolon"},
			{token.IDENT, "x"}, {token.ASSIGN, "="}, {token.IDENT, "x"}, {token.SUB, "-"},
			{token.INT, "1"}, {token.SEMICOLON, "semicolon"}, {token.RBRACE, "}"},
			{token.EOF, "EOF"},
		},
	},

	{
		name:  "Very long identifier",
		input: strings.Repeat("а", 1000),
		tokens: []scannerTestCase{
			{token.IDENT, strings.Repeat("а", 1000)},
			{token.EOF, "EOF"},
		},
	},
	{
		name:  "Very long number",
		input: strings.Repeat("9", 1000),
		tokens: []scannerTestCase{
			{token.INT, strings.Repeat("9", 1000)},
			{token.EOF, "EOF"},
		},
	},
}

func TestScanner(t *testing.T) {
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := scanner.New("test.құрт", []byte(tt.input))
			for _, expected := range tt.tokens {
				sc.Scan()
				tok, lit := sc.Tok(), sc.Lit()
				if tok != expected.tok || lit != expected.lit {
					t.Errorf("%d.%s: got (%v, %s), want (%v, %s)",
						i+1, tt.name, tok, lit, expected.tok, expected.lit)
				}
			}

			sc.Scan()
			if sc.Tok() != token.EOF {
				t.Fatalf("%s: expected EOF after all tokens, got %v", tt.name, sc.Tok())
			}
		})
	}
}
