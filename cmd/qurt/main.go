package main

import (
	"fmt"

	"github.com/nurtai325/qurt/internal/scanner"
	"github.com/nurtai325/qurt/internal/token"
)

func main() {
	input := "hello();hahahah\n12123.&&"
	fmt.Println(input)

	s := scanner.New([]byte(input))

	var tokens []token.Token

	for {
		lit, tok := s.Scan()
		fmt.Println(lit)
		if tok == token.EOF {
			break
		} else if tok == token.ILLEGAL {
			panic(s.Error())
		}

		tokens = append(tokens, tok)
	}

	fmt.Println(tokens)
}
