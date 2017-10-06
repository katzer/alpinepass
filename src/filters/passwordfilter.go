package filters

import "github.com/appPlant/alpinepass/src/data"

// PasswordFilter removes passwords from a Config.
type PasswordFilter struct {
	strings []string
}

func (p PasswordFilter) filter(config data.Config) data.Config {
	config.Password = ""
	for _, u := range config.Users {
		u.Password = ""
	}
	for i := 0; i < len(config.Users); i++ {
		var user = config.Users[i]
		user.Password = ""
		config.Users[i] = user
	}
	return config
}
