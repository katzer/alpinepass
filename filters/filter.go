package filters

import (
	d "github.com/appPlant/alpinepass/data"
	"github.com/urfave/cli"
)

type Filter interface {
	filter(data d.Config) d.Config
}

func FilterConfigs(configs []d.Config, context *cli.Context) []d.Config {
	if !context.GlobalBool("passwords") {
		for i := 0; i < len(configs); i++ {
			configs[i].Password = ""
		}
	}
	return configs
}
