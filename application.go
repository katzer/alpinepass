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

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "show",
			Description: "Print the filtered and unconverted system data to the console.",
			Action: func(c *cli.Context) error {
				return runShowCommand(c)
			},
		},
		cli.Command{
			Name:        "out",
			Description: "Save the system information as a file. The default output format is JSON.",
			Action: func(c *cli.Context) error {
				return runOutCommand(c)
			},
		},
	}

	app.Run(os.Args)
}
