package io

import (
	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/util"
	yaml "gopkg.in/yaml.v2"
)

//ParseYaml creates YAML objects from a string.
func ParseYaml(data string) d.YamlData {
	yamlData := d.YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	util.CheckError(err)
	return yamlData
}
