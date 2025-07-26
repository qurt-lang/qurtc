package scanner

import "errors"

var (
	ErrInvalidCharacter  = errors.New("Рұқсат етілмеген таңба. Жазылған таңбаны тану мүмкін болмады.")
	ErrInvalidIdentifier = errors.New("Рұқсат етілмеген айнымалы немесе функция атауы. Атау әріптен ғана басталып, ары қарай әріптер мен цифрлардан тұру керек. Мысалы: 'атау', 'атау1', 'Атау12', 'АТАУ1', 'h2o'")
	ErrSingleAmpersand   = errors.New("Және операторын қолдану үшін & емес && қолданыңыз.")
	ErrSingleVerticalBar   = errors.New("Немесе операторын қолдану үшін | емес || қолданыңыз.")
)
