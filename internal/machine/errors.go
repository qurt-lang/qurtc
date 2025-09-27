package machine

import "errors"

var (
	ErrNoMain          = errors.New("негізгі деп аталатын функция болуы керек")
	ErrInvalidMain     = errors.New("негізгі функция ешқандай аргумент алмайтын және ештеңе қайтармайтын болуы керек")
	ErrDuplicateStruct = errors.New("бұндай құрылым жарияланып қойған")
	ErrDuplicateFunc   = errors.New("бұндай функция жарияланып қойған")

	ErrCallNoFunc              = errors.New("функция емес мәнді шақыру немесе бұндай функция жоқ")
	ErrNotSameTypeOp           = errors.New("операция тек бірдей типтегі мәндерге қолданылады")
	ErrOpNotSupportedForType   = errors.New("бұл операция мына типке қолданылмайды")
	ErrUnknownOp               = errors.New("бұндай операция жоқ")
	ErrVarExists               = errors.New("бұл атпен айнымалы бар қайтадан жариялай алмайсыз")
	ErrUndefinedReference      = errors.New("бұл атпен айнымалы жоқ")
	ErrFuncArgMismatch         = errors.New("функция шақырылғанда аргументтер дұрыс берілмеген")
	ErrArrAccessOnNotArr       = errors.New("тізім мүшесін алу операциясы тек тізімдерге ғана болады және индекс мәні бүтін шығуы керек")
	ErrStructAccessNotOnStruct = errors.New("құрылым мүшесін алу операциясы тек құрылымдарға ғана болады")
	ErrInvalidAssign           = errors.New("айнымалы мәнін өзгертудің ережелері сақталмаған")
)
