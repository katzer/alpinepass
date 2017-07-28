package app

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"

	"github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/filters"
	"github.com/appPlant/alpinepass/src/io"
	"github.com/appPlant/alpinepass/src/util"
	"github.com/appPlant/alpinepass/src/validation"
)

var context *cli.Context

//execute reads the input, filters it and writes the output.
func execute(appContext *cli.Context) error {
	context = appContext

	configs := io.ReadConfigs(context.GlobalString(FlagInput))
	configs = filters.FilterConfigs(configs, context)

	validate(configs)
	output(configs)

	return nil
}

func validate(configs []data.Config) {
	if !context.GlobalBool(FlagSkip) {
		validation.Validate(configs)
	}
}

func output(configs []data.Config) {
	if context.GlobalBool(FlagDisplay) {
		outputToConsole(configs)
	} else {
		outputToFile(configs)
	}
}

func outputToFile(configs []data.Config) {
	io.WriteJSON(context.GlobalString(FlagOutput), configs, context.GlobalBool(FlagReadable))
}

func outputToConsole(configs []data.Config) {
	var configsJSON []byte
	var err error

	if context.GlobalBool(FlagReadable) {
		configsJSON, err = json.MarshalIndent(configs, "", "    ")
	} else {
		configsJSON, err = json.Marshal(configs)
	}

	util.CheckError(err, "Marshalling JSON failed!")
	fmt.Println(string(configsJSON))
}
