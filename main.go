package main

import (
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
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
		{
			Name:     "help",
			HelpName: "help",
			Action: func(c *cli.Context) error {
				var err error
				if c.NArg() == 0 {
					err = cli.ShowAppHelp(c)
				} else {
					err = cli.ShowCommandHelp(c, c.Args().First())
				}
				return err
			},
			Usage:           `displays help messages.`,
			Description:     `Display help messages.`,
			SkipFlagParsing: true,
			HideHelp:        true,
			HideHelpCommand: true,
		},
	}
}
