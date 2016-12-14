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
			Usage: "Include passwords in the output.",
		},
	}

	app.Action = func(c *cli.Context) error {
		data := readFile()
		yamlData := parseData(data)
		configs := []Config{}

		for _, definition := range yamlData.Defs {
			config := createConfig(definition)
			config = filterConfig(config)
			configs = append(configs, config)

		}

		writeJSON(configs)
		return nil
	}

	app.Run(os.Args)
}
