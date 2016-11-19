package main

import (
	//"flag"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type Reference struct {
}

type Definition struct {
}

// YamlData holds aplinepass' work data
type YamlData struct {
	//map[string][]string
	//map[string]interface{}
	Refs map[string]Reference
	Defs []map[string]string
}

type Config struct {
	configs []ConfigItem
}

type ConfigItem struct {
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

	fmt.Println(yamlData)
	fmt.Printf("%+v\n", yamlData)

	data2, err := yaml.Marshal(&yamlData)
	checkError(err)
	fmt.Println(string(data2))
	writeFile(string(data2), "output.yml")

	conf := ConfigItem{
		Password: "pw1",
		User:     "user1"}
	confJSON, err := json.Marshal(conf)
	checkError(err)
	writeFile(string(confJSON), "output.json")
}
