package filters

import (
	"fmt"

	d "github.com/appPlant/alpinepass/data"
	"github.com/urfave/cli"
)

type Filter interface {
	filter(data d.Config) d.Config
}

func FilterConfigs(configs []d.Config, context *cli.Context) []d.Config {
	if !context.GlobalBool("passwords") {
		passwordFilter := PasswordFilter{}
		configs = applyFilter(configs, passwordFilter)
	}
	if context.GlobalStringSlice("filter") != nil {
		fmt.Println(context.GlobalStringSlice("filter"))
	}
	return configs
}

func applyFilter(configs []d.Config, filter Filter) []d.Config {
	for i := 0; i < len(configs); i++ {
		configs[i] = filter.filter(configs[i])
	}
	return configs
}
