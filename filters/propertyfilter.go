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
	//var remove = false
	fmt.Println(n.Slices)
	return data
}

func filterProperty() {

}
