package filters

import d "github.com/appPlant/alpinepass/data"

//TODO rename this filter and file for better communication about what this filter does
type NameFilter struct {
}

func (p NameFilter) filter(data d.Config) d.Config {
	return data
}
