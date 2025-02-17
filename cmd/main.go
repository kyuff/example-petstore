package main

import (
	"os"

	"github.com/kyuff/anchor"
	"github.com/kyuff/example-petstore/internal/app"
)

func main() {
	os.Exit(app.Run(anchor.DefaultSignalWire()))
}
