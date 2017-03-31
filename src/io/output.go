package io

import (
	"encoding/json"
	"io/ioutil"

	"github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
)

func writeFile(filedata string, filename string) {
	err := ioutil.WriteFile(filename, []byte(filedata), 0644)
	util.CheckError(err, "the file "+filename+"cannot be written!")
}

//WriteJSON writes Config to a file in JSON format.
func WriteJSON(path string, configs []data.Config, pretty bool) {
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
	util.CheckError(err, "Marshalling to JSON failed!")
	writeFile(string(configsJSON), path)
}
