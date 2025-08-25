package parser

import (
	"errors"
	"fmt"

	"github.com/nurtai325/qurt/internal/help"
)

var (
	ErrUnknownDecl     = errors.New("Функция сыртында тек жаңа айнымалы, функция, құрылым жариялауға ғана болады")
	ErrUnexpectedEOF   = errors.New("Файл күтпеген жерден аяқталады")
	ErrUnexpectedToken = errors.New("Күтпеген таңба")
)

func (p *parser) errorAt(err error, helpPage help.DocPage) error {
	errTempl := "Синтаксис қатесі (файл: %s, жол: %d, қатар: %d):\n\t%s\n"

	if helpPage != "" {
		errTempl += fmt.Sprintf("Мына сілтеме сізге қатеңізді түзеуге көмектесуі мүмкін: %s\n", helpPage)
	} else {
		errTempl += fmt.Sprintf("Синтаксис және тілдің ережелері туралы толық ақпарат: %s\n", help.SyntaxPage)
	}

	pos := p.s.Pos()

	return fmt.Errorf(errTempl, pos.File(), pos.Line(), pos.Col(), err.Error())
}
