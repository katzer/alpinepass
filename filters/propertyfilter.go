package filters

import (
	"fmt"
	"strings"

	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/util"
)

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
	for i := 0; i < len(flags); i++ {
		flag := flags[i]
		if !strings.Contains(flag, ":") {
			util.ThrowError("The filter does not contain ':'!")
		}
		if strings.Count(flag, ":") > 1 {
			util.ThrowError("The filter contains too many ':'!")
		}
	}
}
