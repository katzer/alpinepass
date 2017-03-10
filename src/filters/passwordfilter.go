package filters

import d "github.com/appPlant/alpinepass/src/data"

// PasswordFilter removes passwords from a Config.
type PasswordFilter struct {
	strings []string
}

func (p PasswordFilter) filter(data d.Config) d.Config {
	data.Password = ""
	return data
}
