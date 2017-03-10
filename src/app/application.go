package app

import (
	"os"

	"github.com/appPlant/alpinepass/src/util"
	"github.com/urfave/cli"
)

//RunApp sets up the cli application and executes it.
func RunApp() {
	app := cli.NewApp()

	app.Name = "alpinepass"
	app.Version = util.Version
	app.Author = "appPlant GmbH"
	app.Usage = "Manage system environment information."

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "display, d",
			Usage: "Display the output in the console. An output file will not be written. Use this for previewing the generated file.",
		},
		cli.StringSliceFlag{
			Name:  "filter, f",
			Usage: "Filter configurations by name, type and more. Syntax: -f \"[property]:[regex]\"",
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
			Name:  "readable, r",
			Usage: "Make the output more readable.",
		},
		cli.BoolFlag{
			Name:  "skip, s",
			Usage: "Skip the input validation.",
		},
	}

	app.Action = func(context *cli.Context) error {
		return execute(context)
	}

	app.Run(os.Args)
}
