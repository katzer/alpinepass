package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
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

func readConfigs() []Config {
	data := readFile()
	yamlData := parseData(data)
	configs := []Config{}

	for _, definition := range yamlData.Defs {
		config := createConfig(definition)
		configs = append(configs, config)
	}

	return configs
}

func runShowCommand(context *cli.Context) error {
	configs := readConfigs()
	configs = filterConfigs(configs, context)
	configsJSON, err := json.MarshalIndent(configs, "", "    ")
	checkError(err)
	fmt.Println(string(configsJSON))

	return nil
}

func runOutCommand(context *cli.Context) error {
	configs := readConfigs()
	configs = filterConfigs(configs, context)
	writeJSON(configs)

	return nil
}
