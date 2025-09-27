package types

import "errors"

var (
	ErrOutOfBound  = errors.New("тізім ұзындығынан тең немесе одан асатын немесе теріс индекс берілген")
	ErrNotSameType = errors.New("айнымалыға мән бергенде немесе тізімді немесе құрылымды өзгерткенде өзгеретін мүше мен жаңа мәннің типтері бірдей болуы керек")
	ErrNoSuchField = errors.New("бұндай мүше бұл құрылымда жоқ")
	ErrUnknownType = errors.New("бұндай тип жоқ")
)
