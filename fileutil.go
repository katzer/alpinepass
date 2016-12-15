package main

import (
	"encoding/json"
	"io/ioutil"

	d "github.com/appPlant/alpinepass/data"
)

func readFile() string {
	data, err := ioutil.ReadFile("input.yml")
	checkError(err)
	return string(data)
}

func writeFile(data string, filename string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	checkError(err)
}

func writeJSON(configs []d.Config) {
	configsJSON, err := json.MarshalIndent(configs, "", "    ")
	checkError(err)
	writeFile(string(configsJSON), "output.json")
}
