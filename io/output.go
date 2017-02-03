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
func WriteJSON(path string, configs []d.Config) {
	if path == "" {
		path = "output.json"
	}

	configsJSON, err := json.MarshalIndent(configs, "", "    ")
	util.CheckError(err)
	writeFile(string(configsJSON), path)
}
