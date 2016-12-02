package main

import (
	//"flag"

	"github.com/ghodss/yaml"
	//"github.com/davecgh/go-spew/spew"
	"strings"
)

// Separator separates the different parts of an ID
const Separator string = "."

// Filler is used to replace whitespace in strings
const Filler string = "-"

// Definition stores information about a system, used for importing data.
type Definition struct {
	Title    string
	Type     string
	Env      string
	Location string
	User     string
	Password string
	URL      string
	Notes    string
	Tags     []string
}

// YamlData stores information about all systems, used for importing data.
type YamlData struct {
	Defs []Definition
}

// Config stores data about a system, used for exporting data.
type Config struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host,omitempty"`
}

func parseData(data string) YamlData {
	yamlData := YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	checkError(err)
	return yamlData
}

func createConfig(data Definition) Config {
	config := Config{}
	config.Title = data.Title
	config.ID = createID(data)
	config.Password = data.Password
	config.User = data.User
	//TODO host
	return config
}

func createID(data Definition) string {
	id := ""
	id = id + cleanString(data.Location) + Separator
	id = id + cleanString(data.Env) + Separator
	id = id + cleanString(data.Type) + Separator
	id = id + cleanString(data.Title)
	return id
}

func cleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}

func filterConfig(config Config) Config {
	return config
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data := readFile()
	yamlData := parseData(data)
	configs := []Config{}

	for _, definition := range yamlData.Defs {
		config := createConfig(definition)
		config = filterConfig(config)
		configs = append(configs, config)
	}

	writeJSON(configs)
}
