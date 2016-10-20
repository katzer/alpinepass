package main

import (
	//"flag"
	"fmt"
	"io/ioutil"
	//"github.com/go-yaml/yaml"
)

func readFile() string {
	fmt.Printf("Readin input.yaml!\n")
	data, err := ioutil.ReadFile("input.yaml")
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

func writeFile() {

}

func main() {
	data := readFile()
	fmt.Print(data)
}
