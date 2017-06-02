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

		message = verifyCommonFields(config)
		handleMessage(config, message)

		if config.Type == data.DB {
			message = verifyDbFields(config)
			handleMessage(config, message)
		}

		if config.Type == data.SERVER {
			message = verifyServerFields(config)
			handleMessage(config, message)
		}

		if config.Type == data.WEB {
			message = verifyWebFields(config)
			handleMessage(config, message)
		}
	}
	//TODO IDs should be unique!
}

func verifyType(config data.Config) string {
	if !util.StringInArray(config.Type, []string{data.DB, data.SERVER, data.WEB}) {
		return "The Config has an invalid type!"
	}
	return ""
}

//verifyMandatoryFields checks for mandatory fields and returns an error message.
func verifyCommonFields(config data.Config) string {
	if config.ID == "" {
		return "The Config has no ID!"
	}
	if config.Env == "" {
		return "The Config has no Env!"
	}
	if config.Type == "" {
		return "The Config has no Type!"
	}
	if config.URL == "" {
		return "The Config has no URL!"
	}
	return ""
}

func verifyDbFields(config data.Config) string {
	if config.User == "" {
		return "The Config has no User!"
	}
	if config.Server == "" {
		return "The Config has no Server!"
	}
	if config.DB == "" {
		return "The Config has no DB!"
	}
	if config.SID == "" {
		return "The Config has no SID!"
	}
	if config.Host == "" {
		return "The Config has no Host!"
	}
	if config.Port == "" {
		return "The Config has no Port!"
	}
	return ""
}

func verifyServerFields(config data.Config) string {
	if config.User == "" {
		return "The Config has no User!"
	}
	return ""
}

func verifyWebFields(config data.Config) string {
	return ""
}

func handleMessage(config data.Config, message string) {
	if message != "" {
		util.ThrowConfigError(config, message)
	}
}
