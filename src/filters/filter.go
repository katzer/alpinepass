package filters

import (
	"github.com/appPlant/alpinepass/src/data"
	"github.com/urfave/cli"
)

//Filter declares all methods used for implementing filters.
//A Filter takes a Config, modifies its properties and returns it.
type Filter interface {
	filter(config data.Config) data.Config
}

//FilterConfigs applies all filters to each Config.
func FilterConfigs(configs []data.Config, context *cli.Context) []data.Config {
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

//applyFilter
func applyFilter(configs []data.Config, filter Filter) []data.Config {
	cleanedConfigs := []data.Config{}

	for i := 0; i < len(configs); i++ {
		config := filter.filter(configs[i])
		if config.Meta.IsValid {
			cleanedConfigs = append(cleanedConfigs, config)
		}
	}

	return cleanedConfigs
}
