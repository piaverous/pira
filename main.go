package main

import (
	"fmt"
	"os"

	"github.com/piaverous/pira/cmd"
	"github.com/piaverous/pira/pira"
)

func main() {
	app, err := pira.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err)
		os.Exit(1)
	}

	if err := cmd.New(app).Execute(); err != nil {
		os.Exit(1)
	}
}
