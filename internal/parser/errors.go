package parser

import (
	"errors"
	"fmt"

	"github.com/nurtai325/qurtc/internal/help"
)

var (
	ErrUnknownDecl     = errors.New("функция сыртында тек жаңа айнымалы, функция, құрылым жариялауға ғана болады")
	ErrUnexpectedEOF   = errors.New("файл күтпеген жерден аяқталады")
	ErrUnexpectedToken = errors.New("күтпеген таңба немесе cөз")

	ErrInvalidFuncDecl   = errors.New("функция жариялаудың ережелері сақталмаған")
	ErrInvalidStructDecl = errors.New("құрылым жариялаудың ережелері сақталмаған")
	ErrInvalidVarDecl    = errors.New("айнымалы жариялаудың ережелері сақталмаған")
)

func (p *parser) errorAt(err error, helpPage help.DocPage) error {
	// TODO: add source context line
	errTempl := "Синтаксис қатесі (файл: %s, жол: %d, қатар: %d):\n\t%s\n\n"

	if helpPage != "" {
		errTempl += fmt.Sprintf("Мына сілтеме сізге қатеңізді түзеуге көмектесуі мүмкін: %s\n", helpPage)
	} else {
		errTempl += fmt.Sprintf("Синтаксис және тілдің ережелері туралы толық ақпарат: %s\n", help.SyntaxPage)
	}

	pos := p.s.Pos()
	return fmt.Errorf(errTempl, pos.File(), pos.Line(), pos.Col(), err.Error())
}
