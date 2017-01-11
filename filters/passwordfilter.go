package filters

import d "github.com/appPlant/alpinepass/data"

// PasswordFilter removes passwords from a Config.
type PasswordFilter struct {
}

func (p PasswordFilter) filter(data d.Config) d.Config {
	data.Password = ""
	return data
}
