package main

import (
	"os"

	"github.com/urfave/cli"
)

func runApp() {
	app := cli.NewApp()

	app.Name = "alpinepass"
	app.Version = "0.0.0"
	app.Author = "appPlant GmbH"
	app.Usage = "Manage system environment information."

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "filter, f",
			Usage: "Filter configurations by name, type and more.",
		},
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Specify the input file path.",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Specify the output format.",
		},
		cli.BoolFlag{
			Name:  "passwords, p",
			Usage: "Include passwords in the output.",
		},
		cli.BoolFlag{
			Name:  "show, s",
			Usage: "Show the output in the console. An output file will not be written.",
		},
	}

	app.Action = func(context *cli.Context) error {
		if context.GlobalBool("show") {
			return runShowCommand(context)
		} else {
			return runOutCommand(context)
		}
	}

	app.Run(os.Args)
}
