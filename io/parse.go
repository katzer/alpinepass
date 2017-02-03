package io

import (
	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/util"
	yaml "gopkg.in/yaml.v2"
)

// Separator separates the different parts of an ID.
const Separator string = "."

func parseYaml(data string) d.YamlData {
	yamlData := d.YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	util.CheckError(err)
	return yamlData
}

func createID(definition d.Definition) string {
	id := ""
	id = id + util.CleanString(definition.Location) + Separator
	id = id + util.CleanString(definition.Env) + Separator
	id = id + util.CleanString(definition.Type) + Separator
	id = id + util.CleanString(definition.Title)
	return id
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

func ReadConfigs() []d.Config {
	data := readFile()
	yamlData := parseYaml(data)
	configs := []d.Config{}

	for _, definition := range yamlData.Defs {
		config := createConfig(definition)
		configs = append(configs, config)
	}

	return configs
}
