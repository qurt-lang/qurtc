package scanner

import (
	"unicode"
	"unicode/utf8"

	"github.com/nurtai325/qurt/internal/token"
)

type scanner struct {
	src    []byte
	err    error
	cursor int
	line   int
	col    int
}

func New(src []byte) *scanner {
	s := scanner{
		// TODO: use io.Reader and buffer instead of []byte
		src: src,
	}
	return &s
}

func (s *scanner) Pos() (int, int) {
	return s.line, s.col
}

func (s *scanner) Scan() (string, token.Token) {
	ch, chw := s.nextCh()

	if unicode.IsLetter(ch) {
		s.back(chw)
		return s.ident()
	}

	switch ch {

	case -1:
		return "", token.EOF

	case '\n':
		s.line++
		s.col = 0
		return "", token.SEMICOLON

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		s.back(chw)
		return s.numberLit()
	case '"':
		s.back(chw)
		return s.stringLit()

	case '+':
		return "", token.ADD
	case '-':
		return "", token.SUB
	case '*':
		return "", token.MUL
	case '/':
		return "", token.DIV
	case '%':
		return "", token.MOD
	case '&':
		return "", token.LAND
	case '|':
		return "", token.LOR

	case '=':
		ch, chw := s.nextCh()
		if ch == '=' {
			return "", token.EQL
		}
		s.back(chw)
		return "", token.ASSIGN
	case '<':
		ch, chw := s.nextCh()
		if ch == '=' {
			return "", token.LEQ
		}
		s.back(chw)
		return "", token.LSS
	case '>':
		ch, chw := s.nextCh()
		if ch == '=' {
			return "", token.GTR
		}
		s.back(chw)
		return "", token.GEQ
	case '!':
		ch, chw := s.nextCh()
		if ch == '=' {
			return "", token.NEQ
		}
		s.back(chw)
		return "", token.NOT

	case '(':
		return "", token.LPAREN
	case '[':
		return "", token.LBRACK
	case '{':
		return "", token.LBRACE
	case ',':
		return "", token.COMMA
	case '.':
		return "", token.PERIOD
	case ')':
		return "", token.RPAREN
	case ']':
		return "", token.RBRACK
	case '}':
		return "", token.RBRACE
	case ':':
		return "", token.COLON
	case ';':
		return "", token.SEMICOLON
	default:
		s.err = ErrInvalidCharacter
		return "", token.ILLEGAL
	}
}

func (s *scanner) nextCh() (rune, int) {
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

func (s *scanner) ident() (string, token.Token) {
	lit := ""

	for {
		ch, _ := s.nextCh()

		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			lit += string(ch)
			continue
		}

		break
	}

	return lit, token.IDENT
}

func (s *scanner) numberLit() (string, token.Token) {
	lit := ""
	dotSeen := false

	for {
		ch, _ := s.nextCh()

		if unicode.IsDigit(ch) || (!dotSeen && ch == '.') {
			lit += string(ch)
			continue
		}

		break
	}

	if dotSeen {
		return lit, token.FLOAT
	} else {
		return lit, token.INT
	}
}

func (s *scanner) stringLit() (string, token.Token) {
	// skip first '"'
	_, _ = s.nextCh()

	lit := ""

	for {
		ch, _ := s.nextCh()

		if ch == '"' {
			break
		} else if ch == '\\' {
			ch, _ = s.nextCh()
		}

		lit += string(ch)
	}

	return lit, token.STRING
}

func (s *scanner) Error() error {
	return s.err
}
