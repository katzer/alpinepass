package main

import (
	//"flag"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	//"github.com/davecgh/go-spew/spew"
)

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

// YamlData holds aplinepass' work data
type YamlData struct {
	//Refs map[string]Reference
	Defs []Definition
}

type Config struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	User     string `json:"user"`
}

func readFile() string {
	fmt.Println("Reading input.yml!")
	data, err := ioutil.ReadFile("input.yml")
	checkError(err)
	return string(data)
}

func parseData(data string) YamlData {
	yamlData := YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	checkError(err)
	return yamlData
}

func createConfig(data Definition) Config {
	config := Config{}
	return config
}

func filter() {

}

func convert() {

}

func writeFile(data string, filename string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	checkError(err)
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
	for i, definition := range yamlData.Defs {
		fmt.Println(i)
		configs = append(configs, createConfig(definition))
	}

	conf := Config{
		Password: "pw1",
		User:     "user1"}

	moreConfigs := []Config{}
	moreConfigs = append(moreConfigs, conf)
	moreConfigsJSON, err := json.Marshal(moreConfigs)
	checkError(err)
	writeFile(string(moreConfigsJSON), "output.json")
}
