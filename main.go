package main

import (
	"log"
	"os"

	"github.com/arithran/go-cookiecutter/cookiecutter"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cookiecutter",
		Usage: "A command-line utility that creates projects from templates",
		Commands: []*cli.Command{
			{
				Name:      "run",
				Usage:     "cookiecutter run [directory] [find1 replace1 find2 replace2...]",
				UsageText: "cookiecutter run ./templateDir find1 replace1 find2 replace2",
				Action: func(c *cli.Context) error {
					cc, err := cookiecutter.New(c.Args().First(), c.Args().Tail()...)
					if err != nil {
						return err
					}
					return cc.Run()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
