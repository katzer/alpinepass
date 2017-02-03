package app

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"

	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/filters"
	"github.com/appPlant/alpinepass/io"
	"github.com/appPlant/alpinepass/util"
)

// Separator separates the different parts of an ID.
const Separator string = "."

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
	data := io.ReadFile()
	yamlData := io.ParseYaml(data)
	configs := []d.Config{}

	for _, definition := range yamlData.Defs {
		config := createConfig(definition)
		configs = append(configs, config)
	}

	return configs
}

//execute reads the input and writes the output while processing the "show" flag.
func execute(context *cli.Context) error {
	configs := readConfigs()
	configs = filters.FilterConfigs(configs, context)

	if context.GlobalBool("show") {
		configsJSON, err := json.MarshalIndent(configs, "", "    ")
		util.CheckError(err)
		fmt.Println(string(configsJSON))
	} else {
		io.WriteJSON(configs)
	}

	return nil
}
