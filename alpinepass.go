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
	Password string
}

func readFile() string {
	fmt.Println("Reading input.yml!")
	data, err := ioutil.ReadFile("input.yml")
	checkError(err)
	return string(data)
}

func parse(data string) {

}

func filter() {

}

func convert() {

}

func writeFile(data string) {
	err := ioutil.WriteFile("output.yml", []byte(data), 0644)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data := readFile()
	fmt.Print(data)

	yamlData := YamlData{}

	err2 := yaml.Unmarshal([]byte(data), &yamlData)
	checkError(err2)
	fmt.Println(yamlData)
	fmt.Printf("%+v\n", yamlData)

	data2, err := yaml.Marshal(&yamlData)
	checkError(err)
	fmt.Println(string(data2))

	writeFile(data)

	conf := &ConfigItem{
		Password: "1"}
	confJSON, _ := json.Marshal(conf)
	fmt.Println(string(confJSON))
	/*
	 */
}
