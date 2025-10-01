//go:build js

package main

import (
	"strings"
	"syscall/js"

	"github.com/nurtai325/qurtc/internal/exec"
)

const (
	execFnName = "qurtExec"
	filename   = "негізгі.құрт"
)

func Main() error {
	// TODO: try not to panic it stops the whole program in the browser
	js.Global().Set(execFnName, js.FuncOf(func(this js.Value, args []js.Value) any {
		stdout := strings.Builder{}
		source := args[0].String()
		err := exec.Exec(&stdout, filename, []byte(source))
		if err != nil {
			stdout.WriteString(err.Error())
			return stdout.String()
		}
		return stdout.String()
	}))
	select{}
}
