package filters

import d "github.com/appPlant/alpinepass/data"
import "strings"
import "fmt"

//PropertyFilter allows filtering a Config for the content of its properties.
type PropertyFilter struct {
	Slices []string
}

func (p PropertyFilter) filter(data d.Config) d.Config {
	verifyFlags(p.Slices) //TODO verify only once and not for every Config
	//var remove = false
	for i := 0; i < len(p.Slices); i++ {
		filterProperty(p.Slices[i], data)
	}
	return data
}

func filterProperty(property string, data d.Config) {
	var split = strings.Split(property, ":")
	var key = split[0]
	var value = split[1]
	fmt.Println(key + " " + value)
}

func verifyFlags(flags []string) {
	//TODO veryfy that input like "type:Prod" is valid!
}
