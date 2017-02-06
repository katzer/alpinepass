package app

import (
	"os"

	"github.com/urfave/cli"
)

//RunApp sets up the cli application and executes it.
func RunApp() {
	app := cli.NewApp()

	app.Name = "alpinepass"
	app.Version = "0.0.0"
	app.Author = "appPlant GmbH"
	app.Usage = "Manage system environment information."

	app.Flags = []cli.Flag{
		/*
			cli.StringFlag{
				Name:  "format, fo",
				Usage: "Specify the output format.",
			},
		*/
		cli.StringSliceFlag{
			Name:  "filter, f",
			Usage: "Filter configurations by name, type and more.",
		},
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Specify the input file, absolute or relative path.",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Specify the output file, absolute or relative path. An existing file will be overwritten.",
		},
		cli.BoolFlag{
			Name:  "passwords, p",
			Usage: "Include passwords in the output.",
		},
		cli.BoolFlag{
			Name:  "show, s",
			Usage: "Show the output in the console. An output file will not be written. Use this for previewing the generated file.",
		},
	}

	app.Action = func(context *cli.Context) error {
		return execute(context)
	}

	app.Run(os.Args)
}
