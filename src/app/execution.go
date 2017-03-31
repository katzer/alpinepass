package app

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"

	"github.com/appPlant/alpinepass/src/filters"
	"github.com/appPlant/alpinepass/src/io"
	"github.com/appPlant/alpinepass/src/util"
	"github.com/appPlant/alpinepass/src/validation"
)

//execute reads the input, filters it and writes the output.
func execute(context *cli.Context) error {
	configs := io.ReadConfigs(context.GlobalString(FlagInput))
	configs = filters.FilterConfigs(configs, context)

	if !context.GlobalBool(FlagSkip) {
		validation.Validate(configs)
	}

	if context.GlobalBool(FlagDisplay) {
		var configsJSON []byte
		var err error
		if context.GlobalBool(FlagReadable) {
			configsJSON, err = json.MarshalIndent(configs, "", "    ")
		} else {
			configsJSON, err = json.Marshal(configs)
		}
		util.CheckError(err, "Marshalling JSON failed!")
		fmt.Println(string(configsJSON))
	} else {
		io.WriteJSON(context.GlobalString(FlagOutput), configs, context.GlobalBool(FlagReadable))
	}

	return nil
}
