package main

import (
	//"flag"
	"fmt"
	"io/ioutil"
	//"github.com/go-yaml/yaml"
)

func readFile() string {
	fmt.Printf("Readin input.yml!\n")
	data, err := ioutil.ReadFile("input.yml")
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
}

func main() {
	data := readFile()
	fmt.Print(data)

	writeFile(data)
}
