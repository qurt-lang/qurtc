package testutils

import (
	"fmt"
	"os"
)

const examplesDir = "../../examples"

func RunOnExamples(fn func(name string, contents []byte)) error {
	entries, err := os.ReadDir(examplesDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		contents, err := os.ReadFile(fmt.Sprintf("%s/%s", examplesDir, entry.Name()))
		if err != nil {
			return err
		}
		fn(entry.Name(), contents)
	}
	return nil
}
