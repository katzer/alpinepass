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
			Usage: "Include passwords in the output",
		},
	}

	app.Action = func(c *cli.Context) error {
		return runShowCommand(c)
		//return runOutCommand(c)
	}

	app.Run(os.Args)
}
