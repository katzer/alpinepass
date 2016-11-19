package main

import (
	//"flag"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghodss/yaml"
	//"github.com/davecgh/go-spew/spew"
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
	Id       string `json:"id"`
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

func createConfig(data map[string]string) Config {
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

	//fmt.Printf("%+v\n", yamlData.Defs)
	configs := []Config{}
	for i, e := range yamlData.Defs {
		fmt.Println(i)
		spew.Dump(e)
		configs = append(configs, createConfig(e))
	}

	/*
		fmt.Println(yamlData)
		fmt.Printf("%+v\n", yamlData)

		data2, err := yaml.Marshal(&yamlData)
		checkError(err)
		fmt.Println(string(data2))
		writeFile(string(data2), "output.yml")
	*/

	conf := Config{
		Password: "pw1",
		User:     "user1"}

	moreConfigs := []Config{}
	moreConfigs = append(moreConfigs, conf)
	moreConfigsJSON, err := json.Marshal(moreConfigs)
	checkError(err)
	writeFile(string(moreConfigsJSON), "output.json")
}
