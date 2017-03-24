package validation

import (
	"github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
)

//Validate validates the input.
func Validate(configs []data.Config) {
	for i := 0; i < len(configs); i++ {
		config := configs[i]
		message := ""

		message = verifyType(config)
		handleMessage(config, message)

		message = verifyMandatoryFields(config)
		handleMessage(config, message)
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

func verifyType(config data.Config) string {
	if !util.StringInArray(config.Type, []string{"db", "server", "web"}) {
		return "The Config has an invalid type!"
	}
	return ""
}

func handleMessage(config data.Config, message string) {
	if message != "" {
		util.ThrowConfigError(config, message)
	}
}
