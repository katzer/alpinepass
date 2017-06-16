package filters

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
	"github.com/fatih/structs"
)

//ExactSeparator is tha character used for separating filter keys and values for 'exact' matches.
const ExactSeparator string = "="

//ContainsSeparator is tha character used for separating filter keys and values for 'contains' matches.
const ContainsSeparator string = ":"

//PropertyFilter allows filtering a Config for the content of its properties.
type PropertyFilter struct {
	Slices []string
}

func (p PropertyFilter) filter(config data.Config) data.Config {
	verifyFlags(p.Slices) //TODO verify only once and not for every Config
	for i := 0; i < len(p.Slices); i++ {
		config = filterProperty(p.Slices[i], config)
	}
	return config
}

func filterProperty(filter string, config data.Config) data.Config {
	var split = strings.Split(filter, ExactSeparator)
	var key = split[0]
	var value = split[1]
	var regex = regexp.MustCompile(value)

	//TODO Check that key exists in Config field?

	t := reflect.ValueOf(config)

	var field reflect.StructField

	for i := 0; i < t.NumField(); i++ {
		field = t.Type().Field(i)
		if strings.ToLower(field.Name) == key {
			break
		}
	}

	if field.Type.String() == "string" {
		if !regex.MatchString(t.FieldByName(field.Name).String()) {
			config.IsValid = false
		}
	}
	if field.Type.String() == "[]string" {
		var values = t.FieldByName(field.Name)
		var valid = false
		for i := 0; i < values.Len(); i++ {
			if regex.MatchString(values.Index(i).String()) {
				valid = true
				break
			}
		}
		if !valid {
			config.IsValid = false
		}
	}

	return config
}

//verifyFlags checks that the input flags are valid.
func verifyFlags(flags []string) {
	for i := 0; i < len(flags); i++ {
		flag := flags[i]

		//Do some simple checks.
		if !strings.Contains(flag, ExactSeparator) {
			util.ThrowError("The filter does not contain '" + ExactSeparator + "'! Flag: " + flag)
		}
		if strings.Count(flag, ExactSeparator) > 1 {
			util.ThrowError("The filter contains too many '" + ExactSeparator + "'! Flag: " + flag)
		}

		//Check that the key part of a flag matches the fields available in struct Config.
		config := data.Config{}
		fieldNames := structs.Names(config)
		key := strings.ToLower(strings.Split(flag, ExactSeparator)[0])
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
