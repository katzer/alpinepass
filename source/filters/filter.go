package filters

import (
	d "github.com/appPlant/alpinepass/source/data"
	"github.com/urfave/cli"
)

//Filter declares all methods used for implementing filters.
type Filter interface {
	filter(data d.Config) d.Config
}

//FilterConfigs applies all filters to each Config.
func FilterConfigs(configs []d.Config, context *cli.Context) []d.Config {
	if !context.GlobalBool("passwords") {
		passwordFilter := PasswordFilter{}
		configs = applyFilter(configs, passwordFilter)
	}
	if context.GlobalStringSlice("filter") != nil {
		propertyFilter := PropertyFilter{context.GlobalStringSlice("filter")}
		configs = applyFilter(configs, propertyFilter)
	}
	return configs
}

func applyFilter(configs []d.Config, filter Filter) []d.Config {
	cleanedConfigs := []d.Config{}

	for i := 0; i < len(configs); i++ {
		configs[i] = filter.filter(configs[i])
		if configs[i].IsValid {
			cleanedConfigs = append(cleanedConfigs, configs[i])
		}
	}

	return cleanedConfigs
}
