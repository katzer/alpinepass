package validation

import (
	"github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
)

//Validate validates the input.
func Validate(configs []data.Config) {
	//TODO mandatory fields
	for i := 0; i < len(configs); i++ {
		config := configs[i]
		verifyType(config)
		verifyMandatoryFields(config)
	}
	//TODO IDs should be unique!
}

func verifyMandatoryFields(config data.Config) {
	if config.ID == "" {
		util.ThrowConfigError(config, "The Config has no ID!")
	}
	if config.Env == "" {
		util.ThrowConfigError(config, "The Config has no Env!")
	}
}

func verifyType(config data.Config) {
	if !util.StringInArray(config.Type, []string{"db", "server", "web"}) {
		util.ThrowConfigError(config, "The Config has an invalid type!")
	}
}
