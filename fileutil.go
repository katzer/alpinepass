package main

import (
	"encoding/json"
	"io/ioutil"
)

func writeJSON(configs []Config) {
	configsJSON, err := json.MarshalIndent(configs, "", "    ")
	checkError(err)
	writeFile(string(configsJSON), "output.json")
}

func writeFile(data string, filename string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	checkError(err)
}
