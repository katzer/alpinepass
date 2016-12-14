package main

import (
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/urfave/cli"
)

// Separator separates the different parts of an ID
const Separator string = "."

// Filler is used to replace whitespace in strings
const Filler string = "-"

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

func cleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}

func filterConfig(config Config) Config {
	return config
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
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
