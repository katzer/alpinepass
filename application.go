package main

import (
	"fmt"
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

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "show",
			Description: "Print the filtered and unconverted system data to the console.",
			Action: func(c *cli.Context) error {
				fmt.Println("Running the show command!")
				return nil
			},
		},
		cli.Command{
			Name:        "out",
			Description: "Save the system data as a file in the specified output format.",
			Action: func(c *cli.Context) error {
				fmt.Println("Running the out command!")

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
			},
		},
	}

	app.Run(os.Args)
}
