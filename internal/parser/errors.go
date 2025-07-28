package parser

import (
	"errors"
	"fmt"
)

func (p *parser) syntaxError() error {
	return p.newErr("Синтаксис")
}

func (p *parser) lexError() error {
	return p.newErr("Лексика")
}

func (p *parser) newErr(errType string) error {
	errTempl := "%s қатесі (файл: %s, жол: %d, қатар: %d):\n\t%s"
	pos := p.s.Pos()
	return errors.New(fmt.Sprintf(errTempl, errType, pos.File(), pos.Line(), pos.Col(), p.s.Err().Error()))
}
