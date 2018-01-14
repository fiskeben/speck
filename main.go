package main

import (
	"os"

	"github.com/fiskeben/speck/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
