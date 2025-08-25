package scanner

import (
	"unicode"
	"unicode/utf8"

	"github.com/nurtai325/qurt/internal/token"
)

type Scanner interface {
	Scan() bool
	Lit() string
	Tok() token.Token
	Pos() Pos
	Err() error
}

type scanner struct {
	filename string
	src      []byte
	err      error
	cursor   int
	line     int
	col      int
	tok      token.Token
	lit      string
}

func New(filename string, src []byte) Scanner {
	s := scanner{
		filename: filename,
		// TODO: use io.Reader and buffer instead of []byte
		src: src,
	}
	return &s
}

type Pos interface {
	File() string
	Line() int
	Col() int
}

type pos struct {
	filename  string
	line, col int
}

func (p pos) File() string {
	return p.filename
}

func (p pos) Line() int {
	return p.line
}

func (p pos) Col() int {
	return p.col
}

func (s *scanner) Pos() Pos {
	return pos{
		filename: s.filename,
		line:     s.line,
		col:      s.col,
	}
}

func (s *scanner) Lit() string {
	return s.lit
}

func (s *scanner) Tok() token.Token {
	return s.tok
}

func (s *scanner) Scan() bool {
	ch, chw := s.nextCh()

	if ch == '\n' && s.tok != token.SEMICOLON {
		s.lit = "newline"
		s.tok = token.SEMICOLON
		return true
	}

	for unicode.IsSpace(ch) {
		ch, chw = s.nextCh()
	}

	if unicode.IsLetter(ch) {
		s.back(chw)
		s.ident()
		return true
	}

	switch ch {
	case -1:
		s.lit = token.EOF.String()
		s.tok = token.EOF
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		s.back(chw)
		s.numberLit()
	case '"':
		s.back(chw)
		s.stringLit()

	case '+':
		s.lit, s.tok = token.ADD.String(), token.ADD
	case '-':
		s.lit, s.tok = token.SUB.String(), token.SUB
	case '*':
		s.lit, s.tok = token.MUL.String(), token.MUL
	case '/':
		// TODO: add comment consuming and returning comment token here
		s.lit, s.tok = token.DIV.String(), token.DIV
	case '%':
		s.lit, s.tok = token.MOD.String(), token.MOD

	// TODO: all similar repeated code for consuming these two character operators. move to a reusable function
	case '&':
		ch, chw := s.nextCh()
		if ch == '&' {
			s.lit, s.tok = token.LAND.String(), token.LAND
		} else {
			s.back(chw)
			s.err = ErrSingleAmpersand
			s.lit, s.tok = token.ILLEGAL.String(), token.ILLEGAL
		}
	case '|':
		ch, chw := s.nextCh()
		if ch == '|' {
			s.lit, s.tok = token.LOR.String(), token.LOR
		} else {
			s.back(chw)
			s.err = ErrSingleVerticalBar
			s.lit, s.tok = token.ILLEGAL.String(), token.ILLEGAL
		}

	case '=':
		ch, chw := s.nextCh()
		if ch == '=' {
			s.lit, s.tok = token.EQL.String(), token.EQL
		} else {
			s.back(chw)
			s.lit, s.tok = token.ASSIGN.String(), token.ASSIGN
		}

	case '<':
		ch, chw := s.nextCh()
		if ch == '=' {
			s.lit, s.tok = token.LEQ.String(), token.LEQ
		} else {
			s.back(chw)
			s.lit, s.tok = token.LSS.String(), token.LSS
		}

	case '>':
		ch, chw := s.nextCh()
		if ch == '=' {
			s.lit, s.tok = token.GEQ.String(), token.GEQ
		} else {
			s.back(chw)
			s.lit, s.tok = token.GTR.String(), token.GTR
		}

	case '!':
		ch, chw := s.nextCh()
		if ch == '=' {
			s.lit, s.tok = token.NEQ.String(), token.NEQ
		} else {
			s.back(chw)
			s.lit, s.tok = token.NOT.String(), token.NOT
		}

	case '(':
		s.lit, s.tok = token.LPAREN.String(), token.LPAREN
	case '[':
		s.lit, s.tok = token.LBRACK.String(), token.LBRACK
	case '{':
		s.lit, s.tok = token.LBRACE.String(), token.LBRACE
	case ',':
		s.lit, s.tok = token.COMMA.String(), token.COMMA
	case '.':
		s.lit, s.tok = token.PERIOD.String(), token.PERIOD
	case ')':
		s.lit, s.tok = token.RPAREN.String(), token.RPAREN
	case ']':
		s.lit, s.tok = token.RBRACK.String(), token.RBRACK
	case '}':
		s.lit, s.tok = token.RBRACE.String(), token.RBRACE
	case ':':
		s.lit, s.tok = token.COLON.String(), token.COLON
	case ';':
		s.lit, s.tok = "semicolon", token.SEMICOLON

	default:
		s.err = ErrInvalidCharacter
		s.lit, s.tok = token.ILLEGAL.String(), token.ILLEGAL
	}

	if s.tok == token.EOF || s.tok == token.ILLEGAL {
		return false
	} else {
		return true
	}
}

func (s *scanner) nextCh() (rune, int) {
	if s.cursor >= len(s.src) {
		return -1, 0
	}

	s.col += 1

	r, size := utf8.DecodeRune(s.src[s.cursor:])
	if r != utf8.RuneError {
		s.cursor += size
		return r, size
	}

	if size == 0 {
		return -1, 0
	} else {
		return utf8.RuneError, 0
	}
}

func (s *scanner) back(chw int) {
	s.cursor -= chw
}

func (s *scanner) ident() {
	lit := ""

	for {
		ch, chw := s.nextCh()

		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			lit += string(ch)
			continue
		}

		s.back(chw)
		break
	}

	tok, ok := token.Lookup(lit)
	if ok {
		s.lit = lit
		s.tok = tok
	} else {
		s.lit = lit
		s.tok = token.IDENT
	}
}

func (s *scanner) numberLit() {
	lit := ""
	dotSeen := false

	for {
		ch, chw := s.nextCh()

		if unicode.IsDigit(ch) || (!dotSeen && ch == '.') {
			if ch == '.' {
				dotSeen = true
			}

			lit += string(ch)
			continue
		}

		s.back(chw)
		break
	}

	if dotSeen {
		s.lit = lit
		s.tok = token.FLOAT
	} else {
		s.lit = lit
		s.tok = token.INT
	}
}

func (s *scanner) stringLit() {
	// skip first '"'
	_, _ = s.nextCh()

	lit := ""

	for {
		ch, _ := s.nextCh()

		if ch == '"' {
			break
		} else if ch == '\\' {
			lit += string(ch)
			ch, _ = s.nextCh()
		}

		lit += string(ch)
	}

	s.lit = lit
	s.tok = token.STRING
}

func (s *scanner) Err() error {
	return s.err
}
