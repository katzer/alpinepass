package app

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"

	"github.com/appPlant/alpinepass/filters"
	"github.com/appPlant/alpinepass/io"
	"github.com/appPlant/alpinepass/util"
)

//execute reads the input and writes the output while processing the "show" flag.
func execute(context *cli.Context) error {
	configs := io.ReadConfigs()
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
