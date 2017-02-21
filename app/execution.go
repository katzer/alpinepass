package app

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"

	"github.com/appPlant/alpinepass/filters"
	"github.com/appPlant/alpinepass/io"
	"github.com/appPlant/alpinepass/util"
)

//execute reads the input, filters it and writes the output.
func execute(context *cli.Context) error {
	configs := io.ReadConfigs(context.GlobalString("input"))
	configs = filters.FilterConfigs(configs, context)

	if context.GlobalBool("display") {
		var configsJSON []byte
		var err error
		if context.GlobalBool("readable") {
			configsJSON, err = json.MarshalIndent(configs, "", "    ")
		} else {
			configsJSON, err = json.Marshal(configs)
		}
		util.CheckError(err)
		fmt.Println(string(configsJSON))
	} else {
		io.WriteJSON(context.GlobalString("output"), configs, context.GlobalBool("pretty"))
	}

	return nil
}
