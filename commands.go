package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/urfave/cli"
)

// Separator separates the different parts of an ID
const Separator string = "."

func parseData(data string) YamlData {
	yamlData := YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	checkError(err)
	return yamlData
}

func createConfig(definition Definition) Config {
	config := Config{}
	config.Title = definition.Title
	config.ID = createID(definition)
	config.Password = definition.Password
	config.User = definition.User
	//TODO host
	return config
}

func createID(definition Definition) string {
	id := ""
	id = id + cleanString(definition.Location) + Separator
	id = id + cleanString(definition.Env) + Separator
	id = id + cleanString(definition.Type) + Separator
	id = id + cleanString(definition.Title)
	return id
}

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
