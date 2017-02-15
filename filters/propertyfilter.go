package filters

import (
	"fmt"
	"strings"

	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/util"
	"github.com/fatih/structs"
)

//PropertyFilter allows filtering a Config for the content of its properties.
type PropertyFilter struct {
	Slices []string
}

func (p PropertyFilter) filter(data d.Config) d.Config {
	verifyFlags(p.Slices) //TODO verify only once and not for every Config
	for i := 0; i < len(p.Slices); i++ {
		filterProperty(p.Slices[i], data)
	}
	return data
}

func filterProperty(property string, data d.Config) d.Config {
	var split = strings.Split(property, ":")
	var key = split[0]
	var value = split[1]
	fmt.Println(key + " " + value)

	switch key {
	case "location":
		if !strings.Contains(data.Location, value) {
			data.IsValid = false
		}
	}

	return data
}

//verifyFlags checks that the input flags are valid.
func verifyFlags(flags []string) {
	for i := 0; i < len(flags); i++ {
		flag := flags[i]

		//Do some simple checks.
		if !strings.Contains(flag, ":") {
			util.ThrowError("The filter does not contain ':'! Flag: " + flag)
		}
		if strings.Count(flag, ":") > 1 {
			util.ThrowError("The filter contains too many ':'! Flag: " + flag)
		}

		//check that the key part of a flag matches the fields available in struct Config.
		data := d.Config{}
		fieldNames := structs.Names(data)
		key := strings.ToLower(strings.Split(flag, ":")[0])
		isContained := false
		for j := 0; j < len(fieldNames); j++ {
			if key == strings.ToLower(fieldNames[j]) {
				isContained = true
			}
		}
		if !isContained {
			util.ThrowError("The filter type is not valid! Flag: " + flag)
		}
	}
}
