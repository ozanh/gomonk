package main

import (
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ozanh/gomonk/cmd/help"
)

var Version = "v0.1.0"

func main() {
	app := &cli.App{
		Name:     "gomonk",
		Usage:    "Do monitoring with extreme skill.",
		Commands: Commands(os.Stdin),
		Version:  Version,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Commands(reader io.Reader) []*cli.Command {
	return []*cli.Command{
		help.Command(),
	}
}
