//go:build !js

package main

import (
	"errors"
	"os"

	"github.com/nurtai325/qurtc/internal/exec"
)

func Main() error {
	if len(os.Args) != 2 {
		return errors.New("аргумент ретінде код жазылған файл атын беріңіз")
	}
	filename := os.Args[1]
	source, err := os.ReadFile(filename)
	if err != nil {
		return errors.New("берілген атпен файл табылмады")
	}
	return exec.Exec(os.Stdout, filename, source)
}
