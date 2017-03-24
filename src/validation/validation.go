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
		message := verifyMandatoryFields(config)
		if message != "" {
			util.ThrowConfigError(config, "The Config has no ID!")
		}
	}
	//TODO IDs should be unique!
}

//verifyMandatoryFields checks for mandatory fields and returns an error message.
func verifyMandatoryFields(config data.Config) string {
	if config.ID == "" {
		return "The Config has no ID!"
	}
	if config.Env == "" {
		return "The Config has no Env!"
	}

	return ""
}

func verifyType(config data.Config) {
	if !util.StringInArray(config.Type, []string{"db", "server", "web"}) {
		util.ThrowConfigError(config, "The Config has an invalid type!")
	}
}
