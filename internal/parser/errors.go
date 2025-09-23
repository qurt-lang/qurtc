package parser

import (
	"errors"
	"fmt"

	"github.com/nurtai325/qurtc/internal/help"
)

var (
	ErrUnexpectedEOF = errors.New("файл күтпеген жерден аяқталады")

	ErrUnknownDecl     = errors.New("функция сыртында тек жаңа айнымалы, функция, құрылым жариялауға ғана болады")
	ErrInvalidFuncDecl   = errors.New("функция жариялаудың ережелері сақталмаған")
	ErrInvalidStructDecl = errors.New("құрылым жариялаудың ережелері сақталмаған")
	ErrInvalidVarDecl    = errors.New("айнымалы жариялаудың ережелері сақталмаған")

	ErrInvalidIdent    = errors.New("функция, айнымалы, тип атаулары ережеге сай есім болуы керек")
	ErrInvalidArrayLen = errors.New("тізім ұзындығы тек БҮТІН сан ғана бола алады және [] арасында болу керек")
	ErrInvalidTypeName = errors.New("айнымалы немесе функция аргументі типі ережеге сай есім болу керек")
)

func (p *parser) errorAt(err error, helpPage help.DocPage) error {
	// TODO: add source context line
	errTempl := "Синтаксис қатесі (файл: %s, жол: %d, қатар: %d): %s\n"
	if helpPage != "" {
		errTempl += fmt.Sprintf("Мына сілтеме сізге қатеңізді түзеуге көмектесуі мүмкін: %s\n", helpPage)
	} else {
		errTempl += fmt.Sprintf("Синтаксис және тілдің ережелері туралы толық ақпарат: %s\n", help.SyntaxPage)
	}
	pos := p.s.Pos()
	return fmt.Errorf(errTempl, pos.File(), pos.Line(), pos.Col(), err.Error())
}
