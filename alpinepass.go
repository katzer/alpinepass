package main

import (
	"fmt"
	"io/ioutil"
	//"github.com/go-yaml/yaml"
)

func main() {
	fmt.Printf("Readin input.yaml!\n")

	data, err := ioutil.ReadFile("input.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
