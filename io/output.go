package io

import (
	"encoding/json"
	"io/ioutil"

	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/util"
)

func writeFile(data string, filename string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	util.CheckError(err)
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
