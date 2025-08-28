package scanner

import "errors"

var (
	ErrInvalidCharacter  = errors.New("рұқсат етілмеген таңба. Жазылған таңбаны тану мүмкін болмады")
	ErrInvalidIdentifier = errors.New("рұқсат етілмеген айнымалы немесе функция атауы. атау әріптен ғана басталып, ары қарай әріптер мен цифрлардан тұру керек. мысалы: 'атау', 'атау1', 'Атау12', 'АТАУ1', 'h2o'")
	ErrSingleAmpersand   = errors.New("және операторын қолдану үшін & емес && қолданыңыз")
	ErrSingleVerticalBar = errors.New("немесе операторын қолдану үшін | емес || қолданыңыз")
)
