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
