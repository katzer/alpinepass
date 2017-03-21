package io

import (
	"encoding/json"
	"io/ioutil"

	d "github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
)

func writeFile(data string, filename string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	util.CheckError(err)
}

func ConfigToJSON(config d.Config, pretty bool) {

}

func ConfigsToJSON(configs []d.Config, pretty bool) string {
	var configsJSON []byte
	var err error
	if pretty {
		configsJSON, err = json.MarshalIndent(configs, "", "    ")
	} else {
		configsJSON, err = json.Marshal(configs)
	}
	util.CheckError(err)
	return string(configsJSON)
}

//WriteJSON writes Config to a file in JSON format.
func WriteJSON(path string, configs []d.Config, pretty bool) {
	if path == "" {
		path = "output.json"
	}

	var configsJSON []byte
	var err error
	if pretty {
		configsJSON, err = json.MarshalIndent(configs, "", "    ")
	} else {
		configsJSON, err = json.Marshal(configs)
	}
	util.CheckError(err)
	writeFile(string(configsJSON), path)
}
