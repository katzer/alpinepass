package main

import (
	"encoding/json"
	"io/ioutil"

	d "github.com/appPlant/alpinepass/data"
	"github.com/appPlant/alpinepass/util"
)

func readFile() string {
	data, err := ioutil.ReadFile("input.yml")
	util.CheckError(err)
	return string(data)
}

func writeFile(data string, filename string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	util.CheckError(err)
}

func writeJSON(configs []d.Config) {
	configsJSON, err := json.MarshalIndent(configs, "", "    ")
	util.CheckError(err)
	writeFile(string(configsJSON), "output.json")
}
