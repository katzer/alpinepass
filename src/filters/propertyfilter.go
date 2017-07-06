package filters

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
	"github.com/fatih/structs"
)

//ExactSeparator is tha character used for separating filter keys and values for 'exact' matches.
const ExactSeparator string = "="

//ExactRegexTemplate is the regular expression used in combination with ExactSeparator.
const ExactRegexTemplate = "^%s$"

//ContainsSeparator is tha character used for separating filter keys and values for 'contains' matches.
const ContainsSeparator string = ":"

//ContainsRegexTemplate is the regular expression used in combination with ContainsSeparator.
const ContainsRegexTemplate = "%s"

//PropertyFilter allows filtering a Config for the content of its properties.
type PropertyFilter struct {
	Slices []string
}

func (p PropertyFilter) filter(config data.Config) data.Config {
	validateFlags(p.Slices) //TODO verify only once and not for every Config
	for i := 0; i < len(p.Slices); i++ {
		config = filterProperty(p.Slices[i], config)
	}
	return config
}

func filterProperty(filter string, config data.Config) data.Config {
	var separator = getSeparator(filter)
	var split = strings.Split(filter, separator)
	var key = split[0]
	var value = split[1]
	var regex = regexp.MustCompile(fmt.Sprintf(getRegex(separator), value))

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

//validateFlags checks that the input flags are valid.
func validateFlags(flags []string) {
	for i := 0; i < len(flags); i++ {
		flag := flags[i]

		//Check that the filter contains a separator.
		//The "!=" is a hack to emulate XOR which does not exist in
		if !(strings.Contains(flag, ExactSeparator) != strings.Contains(flag, ContainsSeparator)) {
			util.ThrowError("The filter does not contain a correct separator: '" + ExactSeparator + "' or '" + ContainsSeparator + "'\nFaulty flag: " + flag)
		}
		if (strings.Count(flag, ExactSeparator) > 1) != (strings.Count(flag, ContainsSeparator) > 1) {
			util.ThrowError("The filter contains too many '" + ExactSeparator + "'! Flag: " + flag)
		}

		//Check that the key part of a flag matches the fields available in struct Config.
		separator := getSeparator(flag)
		config := data.Config{}
		fieldNames := structs.Names(config)
		key := strings.ToLower(strings.Split(flag, separator)[0])
		isContained := false
		for j := 0; j < len(fieldNames); j++ {
			if key == strings.ToLower(fieldNames[j]) {
				isContained = true
			}
		}
		if !isContained {
			util.ThrowError("The filter type is not valid!\nFaulty flag: " + flag)
		}
	}
}

func getSeparator(filter string) string {
	var separator = ""

	if strings.Contains(filter, ExactSeparator) {
		separator = ExactSeparator
	}
	if strings.Contains(filter, ContainsSeparator) {
		separator = ContainsSeparator
	}

	if separator == "" {
		util.ThrowError("The filter does not contain a valid separator!\nFaulty filter: " + filter)
	}

	return separator
}

// The function does not feel right, maybe a map could be used!
func getRegex(separator string) string {
	if separator == ContainsSeparator {
		return ContainsRegexTemplate
	}
	if separator == ExactSeparator {
		return ExactRegexTemplate
	}

	//TODO Code smell!
	util.ThrowError("There is no regular expression tempate for the given separator!")
	return ""
}
