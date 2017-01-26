package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"

	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/filters"
	"github.com/appPlant/alpinepass/util"
)

// Separator separates the different parts of an ID
const Separator string = "."

func parseData(data string) d.YamlData {
	yamlData := d.YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	util.CheckError(err)
	return yamlData
}

func createConfig(definition d.Definition) d.Config {
	config := d.Config{}
	config.Title = definition.Title
	config.ID = createID(definition)
	config.Password = definition.Password
	config.User = definition.User
	//TODO host
	return config
}

func createID(definition d.Definition) string {
	id := ""
	id = id + util.CleanString(definition.Location) + Separator
	id = id + util.CleanString(definition.Env) + Separator
	id = id + util.CleanString(definition.Type) + Separator
	id = id + util.CleanString(definition.Title)
	return id
}

func readConfigs() []d.Config {
	data := readFile()
	yamlData := parseData(data)
	configs := []d.Config{}

	for _, definition := range yamlData.Defs {
		config := createConfig(definition)
		configs = append(configs, config)
	}

	return configs
}

func runShowCommand(context *cli.Context) error {
	configs := readConfigs()
	configs = filters.FilterConfigs(configs, context)
	configsJSON, err := json.MarshalIndent(configs, "", "    ")
	util.CheckError(err)
	fmt.Println(string(configsJSON))

	return nil
}

func runOutCommand(context *cli.Context) error {
	configs := readConfigs()
	configs = filters.FilterConfigs(configs, context)
	writeJSON(configs)
	return nil
}
