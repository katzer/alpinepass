package app

import (
	"os"

	"github.com/appPlant/alpinepass/src/util"
	"github.com/urfave/cli"
)

const FlagDebug = "debug"
const FlagDisplay = "display"
const FlagFilter = "filter"
const FlagInput = "input"
const FlagOutput = "output"
const FlagPasswords = "passwords"
const FlagReadable = "readable"
const FlagSkip = "skip"

//RunApp sets up the cli application and executes it.
func RunApp() {
	app := cli.NewApp()

	app.Name = "alpinepass"
	app.Version = util.Version
	app.Author = "appPlant GmbH"
	app.Usage = "Manage system environment information."

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  FlagDebug,
			Usage: "Print the stack trace when an error occurs.",
		},
		cli.BoolFlag{
			Name:  FlagDisplay + ", d",
			Usage: "Display the output in the console. An output file will not be written. Use this for previewing the generated file.",
		},
		cli.StringSliceFlag{
			Name:  FlagFilter + ", f",
			Usage: "Filter configurations by name, type and more. Syntax: -f \"[property]:[regex]\"",
		},
		cli.StringFlag{
			Name:  FlagInput + ", i",
			Usage: "Specify the input file, absolute or relative path.",
		},
		cli.StringFlag{
			Name:  FlagOutput + ", o",
			Usage: "Specify the output file, absolute or relative path. An existing file will be overwritten.",
		},
		cli.BoolFlag{
			Name:  FlagPasswords + ", p",
			Usage: "Include passwords in the output.",
		},
		cli.BoolFlag{
			Name:  FlagReadable + ", r",
			Usage: "Make the output more readable.",
		},
		cli.BoolFlag{
			Name:  FlagSkip + ", s",
			Usage: "Skip the input validation.",
		},
	}

	app.Action = func(context *cli.Context) error {
		util.GlobalContext = context
		return execute(context)
	}

	app.Run(os.Args)
}
