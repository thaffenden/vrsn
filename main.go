// Package main is the entrypoint of the CLI.
package main

import (
	"log"

	"github.com/thaffenden/check-version/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
