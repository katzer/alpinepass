package filters

import "github.com/appPlant/alpinepass/src/data"

// PasswordFilter removes passwords from a Config.
type PasswordFilter struct {
	strings []string
}

func (p PasswordFilter) filter(config data.Config) data.Config {
	config.Password = ""
	return config
}
