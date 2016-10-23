package alpinepass

import (
	//"flag"
	"fmt"
	"io/ioutil"
	//"github.com/go-yaml/yaml"
)

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

	writeFile(data)
}
