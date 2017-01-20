package filters

import d "github.com/appPlant/alpinepass/data"
import "fmt"

//TODO rename this filter and file for better communication about what this filter does

//NameFilter allows filtering a Config for name, type and more.
type NameFilter struct {
	Slices []string
}

func (n NameFilter) filter(data d.Config) d.Config {
	fmt.Println(n.Slices)
	return data
}
