package machine

import "errors"

var (
	ErrNoMain          = errors.New("негізгі деп аталатын функция болуы керек")
	ErrInvalidMain     = errors.New("негізгі функция ешқандай аргумент алмайтын және ештеңе қайтармайтын болуы керек")
	ErrDuplicateStruct = errors.New("бұндай құрылым жарияланып қойған")
	ErrDuplicateFunc   = errors.New("бұндай функция жарияланып қойған")

	ErrCallNoFunc = errors.New("функция емес мәнді шақыру немесе бұндай функция жоқ")
)
