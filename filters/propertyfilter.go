package filters

import (
	"fmt"

	d "github.com/appPlant/alpinepass/data"
)

//PropertyFilter allows filtering a Config for the content of its properties.
type PropertyFilter struct {
	Slices []string
}

func (n PropertyFilter) filter(data d.Config) d.Config {
	fmt.Println(n.Slices)
	return data
}
