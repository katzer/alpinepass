package main

import (
	//"flag"
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Reference struct {
	SystemType string
}

// YamlData holds aplinepass' work data
type YamlData struct {
	/*
		Refs struct {
			//map[string][]string
			ProdApp string
			IntApp  string
		}
	*/
	//Refs map[string]interface{}
	SystemTypes map[string]interface{}
	//Defs    string
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

	err := yaml.Unmarshal([]byte(data), &yamlData)
	checkError(err)
	//fmt.Println(yamlData)
	fmt.Printf("%+v\n", yamlData)

	data2, err := yaml.Marshal(&yamlData)
	checkError(err)
	fmt.Println(string(data2))

	writeFile(data)
}
