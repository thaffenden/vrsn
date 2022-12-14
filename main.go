// Package main is the entrypoint of the CLI.
package main

import (
	"fmt"
	"os"

	"github.com/thaffenden/vrsn/cmd"
)

func main() {
	code := 0
	defer func() {
		os.Exit(code)
	}()

	if err := cmd.Execute(); err != nil {
		code = 1
		fmt.Printf("error: %v\n", err)
	}
}
