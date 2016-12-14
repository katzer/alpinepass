package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func runShowCommand(context *cli.Context) error {
	fmt.Println("Running the show command!")
	return nil
}

func runOutCommand(context *cli.Context) error {
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
}
